// Package common @Author:冯铁城 [17615007230@163.com] 2025-07-18 15:03:13
package common

import (
	"os"
	"os/exec"
)

// RunCommand 封装执行命令的函数，传入命令名和参数，返回错误
func RunCommand(name string, args ...string) error {

	//1.创建执行命令对象
	cmd := exec.Command(name, args...)

	//2.标准输出、错误输出重定向到控制台，实时打印
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//3.运行命令，等待完成
	return cmd.Run()
}
