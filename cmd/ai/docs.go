// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

// EmbeddingRecord 文档记录
type EmbeddingRecord struct {
	Id        int64  `json:"id"`
	FileName  string `json:"fileName"`
	FilePath  string `json:"filePath"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// EmbeddingUploadPayload 文档上传请求体
type EmbeddingUploadPayload struct {
	Path string `json:"path"`
}

// EmbeddingUploadResult 文档上传结果
type EmbeddingUploadResult struct {
	AddFiles    []string `json:"addFiles"`
	UpdateFiles []string `json:"updateFiles"`
}

// runListDocs 查看已上传文档
func runListDocs() {

	//1.发送请求
	result, err := doGet(baseURL + "/api/rest/v1/ai/embedding/docs")
	if err != nil {
		fmt.Printf("查询文档失败: %v\n", err)
		return
	}

	//2.解析文档列表
	var docs []EmbeddingRecord
	if err := json.Unmarshal(result.Data, &docs); err != nil {
		fmt.Printf("解析文档列表失败: %v\n", err)
		return
	}

	//3.判空
	if len(docs) == 0 {
		fmt.Println("暂无已上传文档")
		return
	}

	//4.打印文档表格
	fmt.Println("--------------------------------------------------------------------------------")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| ID\t| 文件名\t| 文件路径\t| 创建时间\t| 更新时间\t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	for _, doc := range docs {
		fmt.Fprintf(w, "| %d\t| %s\t| %s\t| %s\t| %s\t|\n",
			doc.Id, doc.FileName, doc.FilePath, doc.CreatedAt, doc.UpdatedAt)
	}
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	w.Flush()
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
