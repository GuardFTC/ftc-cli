// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"fmt"
	"ftcli/common"
)

// runToolsWeb 打开工具管理页面
func runToolsWeb() {

	//1.定义URL
	url := baseURL + "/tools.html"

	//2.打开浏览器
	if err := common.OpenBrowser(url); err != nil {
		fmt.Printf("打开浏览器失败: %v\n", err)
		return
	}

	//3.日志打印
	fmt.Printf("已打开页面: %s\n", url)
}
