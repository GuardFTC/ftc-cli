// Package env @Author:冯铁城 [17615007230@163.com] 2026-06-08 16:00:00
package env

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// runBgLog 滚动查看后台服务日志
func runBgLog(serviceName string) {

	//1.如果未指定服务名，列出可选服务
	if serviceName == "" || serviceName == "list" {
		printAvailableServices()
		return
	}

	//2.查找服务对应的日志文件
	logFile := findServiceLogFile(serviceName)
	if logFile == "" {
		fmt.Printf("未找到服务[%s]的日志配置，可选服务:\n", serviceName)
		printAvailableServices()
		return
	}

	//3.滚动输出日志
	tailLog(logFile)
}

// printAvailableServices 打印可选的后台服务名
func printAvailableServices() {

	//1.获取当前系统对应的项目集合
	systemProjects := envCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		fmt.Println("当前系统无后台服务配置")
		return
	}

	//2.打印提示
	fmt.Println("可选的后台服务:")

	//3.遍历所有项目
	for projectName, projectProperties := range systemProjects {

		//4.遍历项目配置
		for serviceName, propertyValues := range projectProperties {

			//5.只处理后台服务
			if propertyValues[0] != "background" {
				continue
			}

			//6.打印
			fmt.Printf("  * %s (项目: %s)\n", serviceName, projectName)
		}
	}

	//7.打印用法提示
	fmt.Println()
	fmt.Println("用法: ftcli env --bl <服务名>")
}

// findServiceLogFile 根据服务名查找日志文件路径
func findServiceLogFile(serviceName string) string {

	//1.获取当前系统对应的项目集合
	systemProjects := envCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		return ""
	}

	//2.遍历所有项目查找匹配的服务
	for _, projectProperties := range systemProjects {
		for name, propertyValues := range projectProperties {

			//3.只处理后台服务且名称匹配
			if propertyValues[0] == "background" && name == serviceName {
				return propertyValues[1]
			}
		}
	}

	//3.未找到返回空
	return ""
}

// tailLog 滚动输出日志文件（类似 tail -f，先输出最后100行）
func tailLog(logFile string) {

	//1.打开日志文件
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Printf("打开日志文件失败: %v\n", err)
		return
	}
	defer file.Close()

	//2.打印提示
	fmt.Printf(">>> 正在滚动查看日志: %s (按 Ctrl+C 退出)\n", logFile)
	fmt.Println("--------------------------------------------------------------------------------")

	//3.先输出最后100行
	printLastNLines(logFile, 100)

	//4.移动到文件末尾，开始追踪新内容
	file.Seek(0, io.SeekEnd)

	//5.监听退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	//6.创建读取器
	reader := bufio.NewReader(file)

	//7.循环读取新内容
	for {

		//8.检查退出信号
		select {
		case <-sigChan:
			fmt.Println("\n>>> 已退出日志查看")
			return
		default:
		}

		//9.尝试读取一行
		line, err := reader.ReadString('\n')
		if err != nil {

			//10.如果是EOF，等待100ms后继续
			time.Sleep(100 * time.Millisecond)
			continue
		}

		//11.输出日志
		fmt.Print(line)
	}
}

// printLastNLines 打印文件最后N行
func printLastNLines(filePath string, n int) {

	//1.读取文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	//2.按行分割
	lines := strings.Split(string(data), "\n")

	//3.计算起始行
	start := len(lines) - n
	if start < 0 {
		start = 0
	}

	//4.输出
	for _, line := range lines[start:] {
		if line != "" {
			fmt.Println(line)
		}
	}
}
