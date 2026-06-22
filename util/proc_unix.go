//go:build !windows

// Package common @Author:冯铁城 [17615007230@163.com] 2026-06-08 15:00:00
package util

import (
	"os/exec"
	"syscall"
)

// setDetachAttrs 设置进程属性，使子进程脱离终端（Unix下创建新会话）
func setDetachAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}
}
