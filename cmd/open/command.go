// Package open @Author:冯铁城 [17615007230@163.com] 2026-06-08 17:00:00
package open

import (
	"fmt"
	"ftcli/util"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// flag变量
var (
	openList bool
)

// NewOpenCommand 创建open命令
func NewOpenCommand() *cobra.Command {

	//1.设置Flags
	openCmd.Flags().BoolVarP(&openList, "list", "l", false, "输出所有支持的软件")

	//2.返回
	return openCmd
}

// openCmd open命令
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open daily development software",
	Run: func(cmd *cobra.Command, args []string) {

		//1.根据flag执行对应操作
		switch {
		case openList:
			runListApps()
		case len(args) > 0:
			runOpenByName(strings.Join(args, ","))
		default:
			runOpenAlways()
		}
	},
}

// 打印所有支持的软件
func runListApps() {

	//1.获取当前系统对应的软件集合
	apps := openCmdAppsMap[system]
	if apps == nil {
		fmt.Println("当前系统无软件配置")
		return
	}

	//2.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//3.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 软件名称   \t| 默认启动\t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//4.打印软件信息
	for name, config := range apps {
		alwaysFlag := "false"
		if config.Always {
			alwaysFlag = "true"
		}
		fmt.Fprintf(w, "| %s\t| %s\t   |\n", name, alwaysFlag)
	}

	//5.写入控制台
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	w.Flush()
}

// 按名称打开指定软件
func runOpenByName(names string) {

	//1.获取当前系统对应的软件集合
	apps := openCmdAppsMap[system]
	if apps == nil {
		log.Fatalf("当前系统不支持open命令: %v\n", system)
	}

	//2.按逗号分隔软件名称
	nameList := strings.Split(names, ",")

	//3.依次打开软件
	for _, name := range nameList {

		//4.去除空格
		name = strings.TrimSpace(name)

		//5.获取软件配置
		config, exist := apps[name]
		if !exist {
			fmt.Printf(">>> 未找到软件: %s\n", name)
			continue
		}

		//6.打开软件
		openApp(name, config)
	}
}

// 打开所有always为true的软件
func runOpenAlways() {

	//1.获取当前系统对应的软件集合
	apps := openCmdAppsMap[system]
	if apps == nil {
		log.Fatalf("当前系统不支持open命令: %v\n", system)
	}

	//2.遍历打开always为true的软件
	for name, config := range apps {
		if config.Always {
			openApp(name, config)
		}
	}

	//3.打印完成提示
	fmt.Println("所有默认软件已启动！")
}

// 打开单个软件
func openApp(name string, config AppConfig) {

	//1.打印启动提示
	fmt.Printf(">>> 启动 %s\n", name)

	//2.后台启动，不检测存活
	if err := util.RunCommandBackgroundNoCheck(config.Command, config.Args...); err != nil {
		fmt.Printf("    启动失败: %v\n", err)
	}
}
