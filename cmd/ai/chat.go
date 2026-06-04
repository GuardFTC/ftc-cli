// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-03 16:00:00
package ai

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// ChatPayload 聊天请求体
type ChatPayload struct {
	IsLocal     bool   `json:"isLocal"`
	ChatId      string `json:"chatId"`
	UserMessage string `json:"userMessage"`
}

// runChat 进入聊天交互模式
func runChat(isLocal bool) {

	//1.确定聊天模式名称
	modeName := "Web AI 智能检索"
	if isLocal {
		modeName = "Local AI 本地文库"
	}

	//2.获取会话ID
	chatId, err := getChatId()
	if err != nil {
		fmt.Printf("获取会话ID失败: %v\n", err)
		return
	}

	//3.打印进入提示
	fmt.Println("===================================================================")
	fmt.Printf(">> 进入 %s 聊天模式 (输入 'exit' 或 按 Ctrl+C 退出) <<\n", modeName)
	fmt.Println("===================================================================")
	fmt.Println()

	//4.监听Ctrl+C信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//5.创建输入扫描器
	scanner := bufio.NewScanner(os.Stdin)

	//6.聊天循环
	for {

		//7.打印用户输入提示
		fmt.Print("[User] > ")

		//8.监听输入或退出信号
		inputChan := make(chan string, 1)
		go func() {
			if scanner.Scan() {
				inputChan <- scanner.Text()
			} else {
				inputChan <- "exit"
			}
		}()

		//9.等待输入或信号
		var input string
		select {
		case <-sigChan:
			input = "exit"
		case input = <-inputChan:
		}

		//10.处理退出
		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			fmt.Println()
			fmt.Println("===================================================================")
			fmt.Println(">> 已安全退出 AI 聊天模式。感谢使用！")
			fmt.Println("===================================================================")
			return
		}

		//11.空输入跳过
		if input == "" {
			continue
		}

		//12.构建请求
		payload := ChatPayload{
			IsLocal:     isLocal,
			ChatId:      chatId,
			UserMessage: input,
		}

		//13.发送流式请求，打字机效果输出
		fmt.Print("[AI]   > ")
		err := doPostStream(baseURL+"/api/rest/v1/ai/chat/stream", payload, func(chunk string) {
			fmt.Print(cleanMarkdown(chunk))
		})

		//14.处理错误
		if err != nil {
			fmt.Printf("\n请求失败: %v\n", err)
		} else {
			fmt.Println()
		}
		fmt.Println()
	}
}

// getChatId 获取会话ID
func getChatId() (string, error) {

	//1.发送请求
	result, err := doGet(baseURL + "/api/rest/v1/ai/chatId")
	if err != nil {
		return "", err
	}

	//2.解析chatId
	var chatId string
	if err := json.Unmarshal(result.Data, &chatId); err != nil {
		return "", fmt.Errorf("解析会话ID失败: %v", err)
	}

	//3.返回
	return chatId, nil
}
