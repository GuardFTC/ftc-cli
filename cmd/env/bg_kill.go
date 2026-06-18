// Package env @Author:冯铁城 [17615007230@163.com] 2026-06-08 17:30:00
package env

import (
	"fmt"

	"ftcli/common"
)

// runBgKill 停止后台服务
func runBgKill(serviceName string) {

	//1.获取当前系统对应的项目集合
	systemProjects := envCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		fmt.Println("当前系统无后台服务配置")
		return
	}

	//2.查找服务对应的kill配置
	killName, killKeyword := findServiceKillConfig(serviceName)
	if killName == "" {
		fmt.Printf("未找到服务[%s]的配置，用 -b 查看可选服务\n", serviceName)
		return
	}

	//3.执行kill
	fmt.Printf(">>> 停止服务: %s\n", serviceName)
	if err := common.KillProcess(killName, killKeyword); err != nil {
		fmt.Printf("停止失败: %v\n", err)
		return
	}

	//4.打印完成提示
	fmt.Printf(">>> 服务 %s 已停止\n", serviceName)
}

// findServiceKillConfig 根据服务名查找kill配置
func findServiceKillConfig(serviceName string) (string, string) {

	//1.获取当前系统对应的项目集合
	systemProjects := envCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		return "", ""
	}

	//2.遍历所有项目查找匹配的服务
	for _, projectProperties := range systemProjects {
		for name, propertyValues := range projectProperties {

			//3.只处理后台服务且名称匹配
			if propertyValues[0] == "background" && name == serviceName {

				//4.解析kill配置：background, 日志路径, 检测端口, kill进程名, kill关键字, 实际命令...
				killName := propertyValues[3]
				killKeyword := propertyValues[4]
				return killName, killKeyword
			}
		}
	}

	//3.未找到返回空
	return "", ""
}
