// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"github.com/spf13/cobra"
)

// flag变量
var (
	localChat bool
	webChat   bool
	listDocs  bool
	uploadDoc string
	listTools bool
	toolsWeb  bool
	baseURL   string
)

// NewAiCommand 创建ai命令
func NewAiCommand() *cobra.Command {

	//1.设置Flags
	aiCmd.Flags().BoolVarP(&localChat, "local", "l", false, "基于本地文库进行流式AI聊天")
	aiCmd.Flags().BoolVarP(&webChat, "web", "w", false, "基于网络进行流式AI聊天")
	aiCmd.Flags().BoolVarP(&listDocs, "docs", "f", false, "查看已上传文档列表")
	aiCmd.Flags().StringVarP(&uploadDoc, "upload", "u", "", "上传文档(指定文件/目录路径)")
	aiCmd.Flags().BoolVarP(&listTools, "tools", "t", false, "查看AI工具列表")
	aiCmd.Flags().BoolVar(&toolsWeb, "tl", false, "打开AI工具管理页面")
	aiCmd.Flags().StringVarP(&baseURL, "server", "s", defaultBaseURL, "后端服务地址")

	//2.返回
	return aiCmd
}

// aiCmd ai命令
var aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "AI assistant powered by ftcli backend",
	Run: func(cmd *cobra.Command, args []string) {

		//1.根据flag执行对应操作
		switch {
		case localChat:
			runChat(true)
		case webChat:
			runChat(false)
		case listDocs:
			runListDocs()
		case uploadDoc != "":
			runUploadDoc()
		case listTools:
			runListTools()
		case toolsWeb:
			runToolsWeb()
		default:
			cmd.Help()
		}
	},
}
