// Package cmd @Author:冯铁城 [17615007230@163.com] 2025-07-11 14:41:30
package cmd

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"go-ftc-console/common"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "run project env",
	Run: func(cmd *cobra.Command, args []string) {

		//1.如果打印项目列表，则打印并返回，否则执行运行环境命令
		if envListProject {
			consoleEnvProjectInfos()
			return
		} else {
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

	//2.获取项目配置
	properties := systemProjects[envProject]

	//3.创建阻塞器，类似于Java里面的CountDownLatch
	var wg sync.WaitGroup
	wg.Add(len(properties))

	//4.依次执行命令
	for property, propertyValues := range properties {

		//5.创建协程执行命令
		go func(property string, propertyValues []string) {

			//6.命令执行完成后释放阻塞器
			defer wg.Done()

			//7.执行命令
			fmt.Printf(">>> 执行启动%v命令: %v\n", property, propertyValues)
			if err := common.RunCommand(propertyValues[0], propertyValues[1:]...); err != nil {
				fmt.Printf("命令执行失败: %v\n", err)
				return
			}
			fmt.Printf("启动%v命令执行成功\n\n", property)
		}(property, propertyValues)
	}

	//8.阻塞，等待所有协程执行完成
	wg.Wait()

	//9.所有命令完成的提示
	fmt.Println("所有命令执行完成！")
}

// 初始化命令
func initEnv() {

	//1.设置Flags
	envCmd.Flags().StringVarP(&envProject, "project", "p", defaultProject, "项目名称")
	envCmd.Flags().BoolVarP(&envListProject, "list project", "l", false, "输出内置项目信息")

	//2.添加到根命令
	rootCmd.AddCommand(envCmd)
}
