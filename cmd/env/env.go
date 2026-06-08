// Package env Package cmd @Author:冯铁城 [17615007230@163.com] 2025-07-11 14:41:30
package env

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"go-ftc-console/common"

	"github.com/spf13/cobra"
)

// flag变量
var (
	envProject     string
	envListProject bool
	envListBg      bool
	envBgLog       string
)

// NewEnvCommand 创建env命令
func NewEnvCommand() *cobra.Command {

	//1.设置Flags
	envCmd.Flags().StringVarP(&envProject, "project", "p", defaultProject, "项目名称")
	envCmd.Flags().BoolVarP(&envListProject, "list project", "l", false, "输出内置项目信息")
	envCmd.Flags().BoolVarP(&envListBg, "background", "b", false, "查看所有后台运行进程")
	envCmd.Flags().StringVar(&envBgLog, "bl", "", "滚动查看后台服务日志(指定服务名,用-b查看可选服务)")

	//2.返回
	return envCmd
}

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "run project env",
	Run: func(cmd *cobra.Command, args []string) {

		//1.根据flag执行对应操作
		switch {
		case envListProject:
			consoleEnvProjectInfos()
		case envListBg:
			runListBgServices()
		case envBgLog != "":
			runBgLog(envBgLog)
		default:
			runEnvCommand()
		}
	},
}

// 打印项目信息
func consoleEnvProjectInfos() {

	//1.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//2.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 项目名称          \t| 配置        \t| 运行参数                                                         \t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//3.打印项目信息
	for projectName, projectProperties := range envCmdProjectPropertiesMap[system] {

		//4.定义起始行数，起始行数打印项目名称，其他行数只打印项目属性
		line := 1
		for property, propertyValues := range projectProperties {
			if line == 1 {
				fmt.Fprintf(w, "| %-18s\t| %s    \t| %-60s\t|\n", projectName, property, strings.Join(propertyValues, ","))
			} else {
				fmt.Fprintf(w, "| %-18s\t| %s    \t| %-60s\t|\n", "", property, strings.Join(propertyValues, ","))
			}
			line++
		}
		fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	}

	//4.写入控制台
	w.Flush()
}

// 运行环境启动命令
func runEnvCommand() {

	//1.获取当前系统对应的项目集合
	systemProjects := envCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		log.Fatalf("当前系统不支持env命令: %v\n", system)
	}

	//2.获取项目配置
	systemProject := systemProjects[envProject]
	if systemProject == nil {
		log.Fatalf("不支持项目: %v\n", envProject)
	}

	//3.依次执行命令
	for property, propertyValues := range systemProject {

		//4.判断是否为后台启动模式
		if propertyValues[0] == "background" {

			//5.解析配置：background, 日志路径, 检测端口, kill进程名, kill关键字, 实际命令...
			logFile := propertyValues[1]
			checkPort := propertyValues[2]
			killName := propertyValues[3]
			killKeyword := propertyValues[4]
			actualValues := propertyValues[5:]

			//6.先kill旧进程，确保幂等
			fmt.Printf(">>> 停止已运行的%v进程...\n", property)
			common.KillProcess(killName, killKeyword)

			//7.后台启动
			fmt.Printf(">>> 后台启动%v: %v\n", property, actualValues)
			if err := common.RunCommandBackground(logFile, checkPort, actualValues[0], actualValues[1:]...); err != nil {
				fmt.Printf("后台启动失败: %v\n", err)
				continue
			}
		} else {

			//8.前台执行命令
			fmt.Printf(">>> 执行启动%v命令: %v\n", property, propertyValues)
			if err := common.RunCommand(propertyValues[0], propertyValues[1:]...); err != nil {
				fmt.Printf("命令执行失败: %v\n", err)
				continue
			}
			fmt.Printf("启动%v命令执行成功\n\n", property)
		}
	}

	//4.所有命令完成的提示
	fmt.Println("所有命令执行完成！")
}
