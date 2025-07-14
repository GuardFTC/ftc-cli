package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// 定义系统常量
const windows = "windows"
const mac = "darwin"

// 系统名称
var system = runtime.GOOS

// 默认项目
var defaultProject = "prospect-platform"

// 系统-项目名称-项目配置-Map
var projectPropertiesMap = map[string]map[string]map[string]string{
	windows: {
		defaultProject: {
			"pom":    "D:/project/java/prospect-platform/parent/pom.xml",
			"maven":  "D:/base/maven/apache-maven-3.9.9-bin/apache-maven-3.9.9/conf/settings.xml",
			"output": "D:\\project\\java\\prospect-platform\\output",
		},
	},
	mac: {
		defaultProject: {
			"pom":    "",
			"maven":  "",
			"output": "",
		},
	},
}

// flag变量
var (
	project      string
	flagPom      string
	flagMaven    string
	flagOutput   string
	listProjects bool
)

// 打包命令 package java project
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "package java project",
	Run: func(cmd *cobra.Command, args []string) {

		//1.如果打印项目列表，则打印并返回，否则执行打包命令
		if listProjects {
			consoleProjectInfos()
			return
		} else {
			runPackageCommand()
		}
	},
}

// 打印项目信息
func consoleProjectInfos() {

	//1.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//2.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 项目名称          \t| 配置类型 | 路径                                                         \t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//3.打印项目信息
	for projectName, projectProperties := range projectPropertiesMap[system] {
		fmt.Fprintf(w, "| %-18s\t| pom    \t| %-60s\t|\n", projectName, projectProperties["pom"])
		fmt.Fprintf(w, "| %-18s\t| maven  \t| %-60s\t|\n", "", projectProperties["maven"])
		fmt.Fprintf(w, "| %-18s\t| output \t| %-60s\t|\n", "", projectProperties["output"])
		fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	}

	//4.写入控制台
	w.Flush()
}

// 运行打包命令
func runPackageCommand() {

	//1.获取当前系统对应的项目集合
	systemProjects := projectPropertiesMap[system]

	//2.获取项目配置
	pom := ""
	maven := ""
	output := ""
	if systemProject, isExist := systemProjects[project]; isExist {
		pom = systemProject["pom"]
		maven = systemProject["maven"]
		output = systemProject["output"]
	}

	//3.读取自定义项目配置参数
	if flagPom == "" {
		flagPom = pom
	}
	if flagMaven == "" {
		flagMaven = maven
	}
	if flagOutput == "" {
		flagOutput = output
	}

	//4.公共参数，减少代码重复
	baseArgs := []string{
		"-f", flagPom,
		"-s", flagMaven,
		"-DskipTests=true",
	}

	//5.按顺序定义三个命令，clean、install、package
	commands := [][]string{
		append([]string{"clean"}, baseArgs...),
		append([]string{"install"}, baseArgs...),
		append([]string{"package"}, baseArgs...),
	}

	//6.依次执行命令
	for i, args := range commands {

		//7.打印当前执行的命令序号和参数
		fmt.Printf(">>> 执行第%d条命令: mvn %v\n", i+1, args)

		//8.执行命令，出错则打印并退出
		if err := runCommand("mvn", args...); err != nil {
			fmt.Printf("命令执行失败: %v\n", err)
			return
		}

		//9.成功时打印提示
		fmt.Printf("第%d条命令执行成功\n\n", i+1)
	}

	//10.所有命令完成的提示
	fmt.Println("所有命令执行完成！")

	//11.执行完成后，打开目标文件夹
	openOutPutDir()
}

// 封装执行命令的函数，传入命令名和参数，返回错误
func runCommand(name string, args ...string) error {

	//1.创建执行命令对象
	cmd := exec.Command(name, args...)

	//2.标准输出、错误输出重定向到控制台，实时打印
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//3.运行命令，等待完成
	return cmd.Run()
}

// 打开目标文件夹
func openOutPutDir() {

	//1.定义异常
	var err error

	//2.根据不同系统执行不同命令
	if system == windows {
		err = exec.Command("explorer", flagOutput).Start()
	} else {
		err = exec.Command("open", flagOutput).Start()
	}

	//3.判空打印
	if err != nil {
		fmt.Printf("打开文件夹失败: %v\n", err)
	}
}

// 初始化命令
func initPackage() {

	//1.设置Flags
	packageCmd.Flags().StringVarP(&project, "project", "p", defaultProject, "项目名称（优先使用，如果已在内置列表中，无需指定pom/maven/output）")
	packageCmd.Flags().StringVarP(&flagPom, "pom", "P", "", "pom.xml 路径（当项目未被记录时需手动指定）")
	packageCmd.Flags().StringVarP(&flagMaven, "maven", "m", "", "maven settings.xml 路径（当项目未被记录时需手动指定）")
	packageCmd.Flags().StringVarP(&flagOutput, "output", "o", "", "jar 输出目录（当项目未被记录时需手动指定）")
	packageCmd.Flags().BoolVarP(&listProjects, "list project", "l", false, "输出内置项目信息")

	//2.添加到根命令
	rootCmd.AddCommand(packageCmd)
}
