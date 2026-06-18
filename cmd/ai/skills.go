// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-16 16:00:00
package ai

import (
	"fmt"
	"ftcli/common"
)

// runSkillsWeb 打开技能管理页面
func runSkillsWeb() {

	//1.定义URL
	url := baseURL + "/skills.html"

	//2.打开浏览器
	if err := common.OpenBrowser(url); err != nil {
		fmt.Printf("打开浏览器失败: %v\n", err)
		return
	}

	//3.日志打印
	fmt.Printf("已打开页面: %s\n", url)
}
