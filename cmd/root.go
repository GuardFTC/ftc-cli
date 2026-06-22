// Package cmd @Author:冯铁城 [17615007230@163.com] 2025-07-11 14:41:30
package cmd

import (
	"ftcli/cmd/ai"
	"ftcli/cmd/build"
	"ftcli/cmd/env"
	"ftcli/cmd/monitor"
	"ftcli/cmd/open"
	_package "ftcli/cmd/package"
	"ftcli/cmd/sql"
	"os"

	"github.com/spf13/cobra"
)

// 定义根命令
var rootCmd = &cobra.Command{
	Use:     "ftcli",
	Short:   "the tool for ftc develop work",
	Version: "1.0.0",
}

// Init 初始化
func Init() {

	//1.禁用默认补全命令
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	//2.初始化子命令
	rootCmd.AddCommand(env.NewEnvCommand())
	rootCmd.AddCommand(_package.NewPackageCommand())
	rootCmd.AddCommand(sql.NewSqlCommand())
	rootCmd.AddCommand(ai.NewAiCommand())
	rootCmd.AddCommand(build.NewBuildCommand())
	rootCmd.AddCommand(open.NewOpenCommand())
	rootCmd.AddCommand(monitor.NewMonitorCommand())

	//3.执行根命令
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
