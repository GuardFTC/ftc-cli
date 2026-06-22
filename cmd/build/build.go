// Package build @Author:冯铁城 [17615007230@163.com] 2026-06-08 16:00:00
package build

import (
	"fmt"
	"ftcli/util"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// flag变量
var (
	buildProject     string
	buildType        string
	buildListProject bool
)

// NewBuildCommand 创建build命令
func NewBuildCommand() *cobra.Command {

	//1.设置Flags
	buildCmd.Flags().StringVarP(&buildProject, "project", "p", defaultProject, "项目名称")
	buildCmd.Flags().StringVarP(&buildType, "type", "t", defaultType, "构建类型(java/go/all)")
	buildCmd.Flags().BoolVarP(&buildListProject, "list project", "l", false, "输出内置项目信息")

	//2.返回
	return buildCmd
}

// buildCmd build命令
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build project (kill -> package -> start)",
	Run: func(cmd *cobra.Command, args []string) {

		//1.如果打印项目列表，则打印并返回，否则执行构建命令
		if buildListProject {
			consoleBuildProjectInfos()
			return
		} else {
			runBuildCommand()
		}
	},
}

// 打印项目信息
func consoleBuildProjectInfos() {

	//1.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//2.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 项目名称          \t| 支持类型\t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//3.打印项目信息
	for projectName, projectProperties := range buildCmdProjectPropertiesMap[system] {

		//4.判断支持的构建类型
		types := getSupportedTypes(projectProperties)

		//5.打印
		fmt.Fprintf(w, "| %-18s\t| %s\t|\n", projectName, types)
		fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	}

	//4.写入控制台
	w.Flush()
}

// 获取项目支持的构建类型
func getSupportedTypes(projectProperties map[string][]string) string {

	//1.定义结果
	result := ""

	//2.判断是否支持java
	if _, ok := projectProperties["java-pom"]; ok {
		result += "java "
	}

	//3.判断是否支持go
	if _, ok := projectProperties["go-source"]; ok {
		result += "go "
	}

	//4.返回
	return result
}

// 运行构建命令
func runBuildCommand() {

	//1.获取当前系统对应的项目集合
	systemProjects := buildCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		log.Fatalf("当前系统不支持build命令: %v\n", system)
	}

	//2.获取项目配置
	projectProperties := systemProjects[buildProject]
	if projectProperties == nil {
		log.Fatalf("不支持项目: %v\n", buildProject)
	}

	//3.根据构建类型执行对应操作
	switch buildType {
	case typeJava:
		buildJava(projectProperties)
	case typeGo:
		buildGo(projectProperties)
	case typeAll:
		buildJava(projectProperties)
		buildGo(projectProperties)
	default:
		log.Fatalf("不支持的构建类型: %v，可选: java/go/all\n", buildType)
	}
}

// 构建Java项目
func buildJava(projectProperties map[string][]string) {

	//1.打印开始构建提示
	fmt.Println("===================================================================")
	fmt.Println(">>> [Java] 开始构建")
	fmt.Println("===================================================================")

	//2.杀死已运行的Java进程
	fmt.Println(">>> [Java] 停止已运行的进程...")
	killItems := projectProperties["java-kill"]
	if err := util.KillProcess(killItems[0], killItems[1]); err != nil {
		log.Fatalf("停止Java进程失败: %v\n", err)
	}

	//3.执行Maven打包
	fmt.Println(">>> [Java] 执行Maven打包...")
	if err := runMavenBuild(projectProperties); err != nil {
		log.Fatalf("Maven打包失败: %v\n", err)
	}

	//4.后台启动Java服务
	fmt.Println(">>> [Java] 后台启动服务...")
	logFile := projectProperties["java-log"][0]
	checkPort := projectProperties["java-port"][0]
	startItems := projectProperties["java-start"]
	if err := util.RunCommandBackground(logFile, checkPort, startItems[0], startItems[1:]...); err != nil {
		fmt.Printf("Java服务启动失败: %v\n", err)
		return
	}

	//5.打印完成提示
	fmt.Println(">>> [Java] 构建完成，服务已在后台启动")
	fmt.Println()
}

// 执行Maven打包
func runMavenBuild(projectProperties map[string][]string) error {

	//1.获取Maven配置
	pom := projectProperties["java-pom"][0]
	maven := projectProperties["java-maven"][0]

	//2.定义公共参数
	baseArgs := []string{
		"-f", pom,
		"-s", maven,
		"-DskipTests=true",
	}

	//3.按顺序定义三个命令，clean、install、package
	commands := [][]string{
		append([]string{"clean"}, baseArgs...),
		append([]string{"install"}, baseArgs...),
		append([]string{"package"}, baseArgs...),
	}

	//4.依次执行命令
	for i, args := range commands {

		//5.打印当前执行的命令序号和参数
		fmt.Printf(">>> 执行第%d条命令: mvn %v\n", i+1, args)

		//6.执行命令，出错则打印并退出
		if err := util.RunCommand("mvn", args...); err != nil {
			fmt.Printf("命令执行失败: %v\n", err)
			return err
		}

		//7.成功时打印提示
		fmt.Printf("第%d条命令执行成功\n\n", i+1)
	}

	//8.打印完成提示
	fmt.Println("Maven打包完成！")

	//9.默认返回
	return nil
}

// 构建Go项目
func buildGo(projectProperties map[string][]string) {

	//1.打印开始构建提示
	fmt.Println("===================================================================")
	fmt.Println(">>> [Go] 开始构建")
	fmt.Println("===================================================================")

	//2.获取Go配置
	source := projectProperties["go-source"][0]
	output := projectProperties["go-output"][0]

	//3.根据系统选择不同的构建策略
	if system == windows {
		buildGoWindows(source, output)
	} else {
		buildGoUnix(source, output)
	}

	//4.打印完成提示
	fmt.Println(">>> [Go] 构建完成")
	fmt.Println()
}

// Windows下构建Go项目（先编译到临时文件，再通过bat脚本延迟替换）
func buildGoWindows(source string, output string) {

	//1.计算绝对路径
	outputDir := filepath.Dir(filepath.Join(source, output))
	outputName := filepath.Base(output)
	targetPath := filepath.Join(outputDir, outputName)
	tempPath := filepath.Join(outputDir, "ftcli_new.exe")

	//2.编译到临时文件
	fmt.Printf(">>> [Go] 编译项目: %s -> %s\n", source, tempPath)
	if err := util.RunCommandInDir(source, "go", "build", "-o", tempPath); err != nil {
		log.Fatalf("Go编译失败: %v\n", err)
	}

	//3.生成延迟替换bat脚本
	batPath := filepath.Join(outputDir, "ftcli_replace.bat")
	batContent := fmt.Sprintf("@echo off\r\n:loop\r\ntasklist /FI \"PID eq %d\" 2>NUL | find \"%d\" >NUL\r\nif not errorlevel 1 (\r\n    timeout /t 1 /nobreak >NUL\r\n    goto loop\r\n)\r\nmove /Y \"%s\" \"%s\"\r\ndel \"%%~f0\"\r\n", os.Getpid(), os.Getpid(), tempPath, targetPath)

	//4.写入bat脚本
	if err := os.WriteFile(batPath, []byte(batContent), 0644); err != nil {
		log.Fatalf("生成替换脚本失败: %v\n", err)
	}

	//5.后台启动bat脚本（不需要检测存活）
	fmt.Printf(">>> [Go] 延迟替换脚本已生成，当前进程退出后将自动替换 %s\n", targetPath)
	if err := util.RunCommandBackgroundNoCheck("cmd", "/C", batPath); err != nil {
		log.Fatalf("启动替换脚本失败: %v\n", err)
	}
}

// Unix下构建Go项目（直接覆盖即可）
func buildGoUnix(source string, output string) {

	//1.执行go build
	fmt.Printf(">>> [Go] 编译项目: %s -> %s\n", source, output)
	if err := util.RunCommandInDir(source, "go", "build", "-o", output); err != nil {
		log.Fatalf("Go编译失败: %v\n", err)
	}
}
