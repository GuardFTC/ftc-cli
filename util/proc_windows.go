//go:build windows

// Package common @Author:冯铁城 [17615007230@163.com] 2026-06-08 15:00:00
package util

import (
	"os/exec"
	"syscall"
)

// setDetachAttrs 设置进程属性，使子进程完全独立于父进程（Windows下创建新进程组）
func setDetachAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
}
