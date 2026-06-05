// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"
)

// ToolSpec 工具规格
type ToolSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// runListTools 查看工具列表
func runListTools() {

	//1.发送请求
	result, err := doGet(baseURL + "/api/rest/v1/ai/tools")
	if err != nil {
		fmt.Printf("查询工具失败: %v\n", err)
		return
	}

	//2.解析工具列表
	var tools []ToolSpec
	if err := json.Unmarshal(result.Data, &tools); err != nil {
		fmt.Printf("解析工具列表失败: %v\n", err)
		return
	}

	//3.判空
	if len(tools) == 0 {
		fmt.Println("暂无已注册工具")
		return
	}

	//4.打印工具表格
	fmt.Println("--------------------------------------------------------------------------------")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 工具名称\t| 工具描述\t| 工具类型\t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	for _, tool := range tools {
		fmt.Fprintf(w, "| %s\t| %s\t| %s\t|\n", tool.Name, tool.Description, tool.Type)
	}
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	w.Flush()
}

// runToolsWeb 打开工具管理页面
func runToolsWeb() {

	//1.拼接URL
	url := baseURL + "/tools.html"

	//2.根据系统打开浏览器
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	//3.执行命令
	if err := cmd.Run(); err != nil {
		fmt.Printf("打开浏览器失败: %v\n", err)
		return
	}

	//4.打印提示
	fmt.Printf("已打开工具管理页面: %s\n", url)
}
