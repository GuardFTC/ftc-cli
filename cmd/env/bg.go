// Package env @Author:冯铁城 [17615007230@163.com] 2026-06-08 16:00:00
package env

import (
	"fmt"
	"os"
	"text/tabwriter"

	"go-ftc-console/common"
)

// runListBgServices 查看所有后台运行进程状态
func runListBgServices() {

	//1.获取当前系统对应的项目集合
	systemProjects := envCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		fmt.Println("当前系统无后台服务配置")
		return
	}

	//2.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//3.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 项目\t             | 服务名\t      | 端口\t| 状态\t  | 日志路径\t")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//4.遍历所有项目
	for projectName, projectProperties := range systemProjects {

		//5.遍历项目配置
		for serviceName, propertyValues := range projectProperties {

			//6.只处理后台服务
			if propertyValues[0] != "background" {
				continue
			}

			//7.解析配置
			logFile := propertyValues[1]
			checkPort := propertyValues[2]
			killName := propertyValues[3]
			killKeyword := propertyValues[4]

			//8.检查进程是否存活
			status := "未运行"
			if common.IsProcessRunning(killName, killKeyword) {
				status = "运行中"
			}

			//9.打印
			fmt.Fprintf(w, "| %s\t| %s\t| %s\t| %s\t| %s\t\n",
				projectName, serviceName, checkPort, status, logFile)
		}
	}

	//10.写入控制台
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	w.Flush()
}
