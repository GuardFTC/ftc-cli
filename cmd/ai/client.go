// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// httpClient 全局HTTP客户端
var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

// streamClient 流式HTTP客户端(无超时，由服务端控制)
var streamClient = &http.Client{}

// RestfulResult 统一响应体
type RestfulResult struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// doGet 发送GET请求
func doGet(url string) (*RestfulResult, error) {

	//1.发送请求
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	//2.读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	//3.解析响应
	var result RestfulResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	//4.返回
	return &result, nil
}

// doPost 发送POST请求
func doPost(url string, payload interface{}) (*RestfulResult, error) {

	//1.序列化请求体
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %v", err)
	}

	//2.发送请求
	resp, err := httpClient.Post(url, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	//3.读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	//4.解析响应
	var result RestfulResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	//5.返回
	return &result, nil
}

// doPostStream 发送POST请求并流式读取SSE响应，逐chunk回调
func doPostStream(url string, payload interface{}, onChunk func(string)) error {

	//1.序列化请求体
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("序列化请求体失败: %v", err)
	}

	//2.构建请求
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBytes))
	if err != nil {
		return fmt.Errorf("构建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	//3.发送请求
	resp, err := streamClient.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	//4.校验状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("服务端响应异常[%d]: %s", resp.StatusCode, string(body))
	}

	//5.按SSE协议解析流：同一事件内的多个data行用\n拼接，空行表示事件结束
	scanner := bufio.NewScanner(resp.Body)

	//5.1扩大缓冲区，防止长行截断
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	//5.2当前事件的data行缓冲
	var dataLines []string

	//5.3刷新事件：将累积的data行用\n拼接后回调
	flush := func() {
		if len(dataLines) > 0 {
			onChunk(strings.Join(dataLines, "\n"))
			dataLines = dataLines[:0]
		}
	}

	//5.4逐行读取
	for scanner.Scan() {
		line := scanner.Text()

		//6.空行表示一个SSE事件结束，刷新输出
		if line == "" {
			flush()
			continue
		}

		//7.data行：去掉"data:"前缀后累积(保留原始空格，确保代码缩进不丢失)
		if strings.HasPrefix(line, "data:") {
			dataLines = append(dataLines, line[len("data:"):])
		}
	}

	//8.刷新最后一个未以空行结尾的事件
	flush()

	//7.检查扫描错误
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取流失败: %v", err)
	}

	//8.返回
	return nil
}
