// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"encoding/json"
	"fmt"
	"ftcli/util"
)

// EmbeddingUploadPayload 文档上传请求体
type EmbeddingUploadPayload struct {
	Path string `json:"path"`
}

// EmbeddingUploadResult 文档上传结果
type EmbeddingUploadResult struct {
	AddFiles    []string `json:"addFiles"`
	UpdateFiles []string `json:"updateFiles"`
}

// runUploadDoc 上传文档
func runUploadDoc() {

	//1.构建请求体
	payload := EmbeddingUploadPayload{Path: uploadDoc}

	//2.发送请求
	result, err := doPost(baseURL+"/api/rest/v1/ai/embedding/docs", payload)
	if err != nil {
		fmt.Printf("上传文档失败: %v\n", err)
		return
	}

	//3.解析上传结果
	var uploadResult EmbeddingUploadResult
	if err := json.Unmarshal(result.Data, &uploadResult); err != nil {
		fmt.Printf("解析上传结果失败: %v\n", err)
		return
	}

	//4.打印结果
	fmt.Printf("上传完成！新增文档: %d 个, 更新文档: %d 个\n", len(uploadResult.AddFiles), len(uploadResult.UpdateFiles))
	if len(uploadResult.AddFiles) > 0 {
		fmt.Println("新增文件:")
		for _, f := range uploadResult.AddFiles {
			fmt.Printf("  + %s\n", f)
		}
	}
	if len(uploadResult.UpdateFiles) > 0 {
		fmt.Println("更新文件:")
		for _, f := range uploadResult.UpdateFiles {
			fmt.Printf("  ~ %s\n", f)
		}
	}
}

// runDocsWeb 打开文档管理页面
func runDocsWeb() {

	//1.定义URL
	url := baseURL + "/docs.html"

	//2.打开浏览器
	if err := util.OpenBrowser(url); err != nil {
		fmt.Printf("打开浏览器失败: %v\n", err)
		return
	}

	//3.日志打印
	fmt.Printf("已打开页面: %s\n", url)
}
