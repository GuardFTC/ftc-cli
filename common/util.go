// Package common @Author:冯铁城 [17615007230@163.com] 2025-07-18 15:03:13
package common

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

// RunCommand 封装执行命令的函数，传入命令名和参数，返回错误
func RunCommand(command string, args ...string) error {

	//1.创建执行命令对象
	cmd := exec.Command(command, args...)

	//2.标准输出、错误输出重定向到控制台，实时打印
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//3.运行命令，等待完成
	return cmd.Run()
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

// IsNumeric 判断是否为数字
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
