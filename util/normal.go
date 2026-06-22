// Package common @Author:冯铁城 [17615007230@163.com] 2025-07-18 15:03:13
package util

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

// IsNumeric 判断是否为数字
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// GetProjectItems 获取项目配置项
func GetProjectItems(project map[string][]string, key string) []string {

	//1.获取项目配置项
	projectItems := project[key]

	//2.配置项判空
	if projectItems == nil || len(projectItems) == 0 {
		log.Fatalf("项目不包含配置项:%s\n", key)
	}

	//3.返回
	return projectItems
}

// IsProcessRunning 判断进程是否正在运行
func IsProcessRunning(nameContains string, cmdContains string) bool {

	//1.参数判定
	if nameContains == "" && cmdContains == "" {
		return false
	}

	//2.获取所有进程
	processes, err := process.Processes()
	if err != nil {
		return false
	}

	//3.遍历所有进程
	for _, p := range processes {

		//4.获取进程名称，进程命令行
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()

		//5.如果进程名包含关键字，并且命令行包含关键字
		if strings.Contains(strings.ToLower(name), nameContains) && strings.Contains(cmdline, cmdContains) {
			return true
		}
	}

	//6.默认返回
	return false
}

// KillProcess 杀死进程
func KillProcess(nameContains string, cmdContains string) error {

	//1.参数判定
	if nameContains == "" && cmdContains == "" {
		return errors.New("参数错误")
	}

	//2.获取所有进程
	processes, err := process.Processes()
	if err != nil {
		return err
	}

	//3.遍历所有进程
	for _, p := range processes {

		//5.获取进程名称，进程命令行
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()

		//4.如果进程名包含关键字，并且命令行包含关键字
		if strings.Contains(strings.ToLower(name), nameContains) && strings.Contains(cmdline, cmdContains) {

			//5.杀死进程
			fmt.Printf("准备杀死进程: PID=%d, Name=%s, Cmdline=%s\n", p.Pid, name, cmdline)
			if err := p.Kill(); err != nil {
				fmt.Printf("杀死进程失败: %v\n", err)
				return err
			} else {
				fmt.Printf("成功杀死进程: PID=%d\n", p.Pid)
			}
		}
	}

	//6.默认返回
	return nil
}

// OpenBrowser 用默认浏览器打开URL
func OpenBrowser(url string) error {

	//1.定义系统命令
	var cmd *exec.Cmd

	//2.根据不同系统选择不同命令
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	//3.执行命令
	return cmd.Run()
}
