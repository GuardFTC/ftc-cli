// Package common @Author:冯铁城 [17615007230@163.com] 2026-06-08 14:50:54
package util

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

// utf8Once 确保UTF-8代码页设置只执行一次
var utf8Once sync.Once

// RunCommand 封装执行命令的函数，传入命令名和参数，返回错误
func RunCommand(command string, args ...string) error {

	//1.确保控制台使用UTF-8代码页
	ensureUTF8Console()

	//2.创建执行命令对象
	cmd := exec.Command(command, args...)

	//3.标准输出、错误输出重定向到控制台，实时打印
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//4.运行命令，等待完成
	return cmd.Run()
}

// RunCommandInDir 在指定目录下执行命令，传入工作目录、命令名和参数，返回错误
func RunCommandInDir(dir string, command string, args ...string) error {

	//1.确保控制台使用UTF-8代码页
	ensureUTF8Console()

	//2.创建执行命令对象
	cmd := exec.Command(command, args...)

	//3.设置工作目录
	cmd.Dir = dir

	//4.标准输出、错误输出重定向到控制台，实时打印
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//5.运行命令，等待完成
	return cmd.Run()
}

// RunCommandBackground 后台启动进程，日志输出到指定文件，可选端口检测，返回错误
func RunCommandBackground(logFile string, checkPort string, command string, args ...string) error {

	//1.确保控制台使用UTF-8代码页
	ensureUTF8Console()

	//2.创建执行命令对象
	cmd := exec.Command(command, args...)

	//3.设置进程属性，使子进程完全独立于父进程（Windows下创建新进程组）
	setDetachAttrs(cmd)

	//4.如果指定了日志文件，则将stdout和stderr重定向到文件
	if logFile != "" {

		//5.确保日志目录存在
		logDir := filepath.Dir(logFile)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return fmt.Errorf("创建日志目录失败: %v", err)
		}

		//6.打开日志文件（追加模式）
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("打开日志文件失败: %v", err)
		}

		//7.重定向stdout和stderr到日志文件
		cmd.Stdout = file
		cmd.Stderr = file
	} else {

		//8.不指定日志文件时，丢弃输出
		cmd.Stdout = nil
		cmd.Stderr = nil
	}

	//9.不绑定标准输入
	cmd.Stdin = nil

	//10.启动进程，不等待结束
	if err := cmd.Start(); err != nil {
		return err
	}

	//11.打印进程PID和日志路径
	fmt.Printf(">>> 服务已后台启动，PID=%d\n", cmd.Process.Pid)
	if logFile != "" {
		fmt.Printf(">>> 日志文件: %s\n", logFile)
	}

	//12.等待进程启动，通过端口检测是否成功（最多等待15秒）
	pid := cmd.Process.Pid
	go func() {
		cmd.Wait()
	}()

	//13.如果指定了检测端口，通过端口判断启动是否成功
	if checkPort != "" {
		if waitForPort(checkPort, 15*time.Second) {
			fmt.Printf(">>> 服务运行中，PID=%d，端口%s已就绪\n", pid, checkPort)
		} else {
			fmt.Printf(">>> [警告] 等待端口%s超时，可能启动失败！\n", checkPort)
			if logFile != "" {
				printLogTail(logFile, 5)
			}
			return fmt.Errorf("后台进程启动后未能监听端口%s", checkPort)
		}
	} else {

		//14.未指定端口时，等待3秒后检查进程是否存活
		time.Sleep(3 * time.Second)
		if !isProcessAlive(pid) {
			fmt.Printf(">>> [警告] 进程PID=%d已退出，可能启动失败！\n", pid)
			if logFile != "" {
				printLogTail(logFile, 5)
			}
			return fmt.Errorf("后台进程启动后立即退出")
		}
		fmt.Printf(">>> 服务运行中，PID=%d\n", pid)
	}

	//15.默认返回
	return nil
}

// RunCommandBackgroundNoCheck 后台启动进程，不做任何存活检测，返回错误
func RunCommandBackgroundNoCheck(command string, args ...string) error {

	//1.创建执行命令对象
	cmd := exec.Command(command, args...)

	//2.设置进程属性，使子进程完全独立于父进程
	setDetachAttrs(cmd)

	//3.丢弃子进程的标准输出和错误输出，防止污染父进程终端
	devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err == nil {
		cmd.Stdout = devNull
		cmd.Stderr = devNull
		defer devNull.Close()
	}
	cmd.Stdin = nil

	//4.启动进程，不等待结束
	if err := cmd.Start(); err != nil {
		return err
	}

	//5.释放进程资源
	go cmd.Wait()

	//6.默认返回
	return nil
}

// ensureUTF8Console 确保Windows控制台使用UTF-8代码页，解决中文乱码问题（仅执行一次）
func ensureUTF8Console() {
	if runtime.GOOS == "windows" {
		utf8Once.Do(func() {
			exec.Command("cmd", "/C", "chcp", "65001").Run()
		})
	}
}

// waitForPort 等待端口就绪，每秒检测一次
func waitForPort(port string, timeout time.Duration) bool {

	//1.计算截止时间
	deadline := time.Now().Add(timeout)

	//2.循环检测
	for time.Now().Before(deadline) {

		//3.尝试连接端口
		conn, err := net.DialTimeout("tcp", "127.0.0.1:"+port, 1*time.Second)
		if err == nil {
			conn.Close()
			return true
		}

		//4.等待1秒后重试
		time.Sleep(1 * time.Second)
	}

	//5.超时返回
	return false
}

// isProcessAlive 检查进程是否存活
func isProcessAlive(pid int) bool {

	//1.查找进程
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return false
	}

	//2.获取进程状态
	status, err := p.Status()
	if err != nil {
		return false
	}

	//3.判断是否存活
	return len(status) > 0
}

// printLogTail 打印日志文件最后几行
func printLogTail(logFile string, lines int) {

	//1.读取日志文件
	data, err := os.ReadFile(logFile)
	if err != nil {
		return
	}

	//2.按行分割
	allLines := strings.Split(string(data), "\n")

	//3.取最后N行
	start := len(allLines) - lines
	if start < 0 {
		start = 0
	}

	//4.打印
	fmt.Println(">>> 日志文件最后几行:")
	for _, line := range allLines[start:] {
		if strings.TrimSpace(line) != "" {
			fmt.Printf("    %s\n", line)
		}
	}
}
