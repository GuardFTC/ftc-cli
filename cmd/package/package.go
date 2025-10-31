// Package _package @Author:冯铁城 [17615007230@163.com] 2025-07-11 14:41:30
package _package

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"go-ftc-console/common"

	"github.com/spf13/cobra"
)

// NewPackageCommand 创建package命令
func NewPackageCommand() *cobra.Command {

	//1.设置Flags
	packageCmd.Flags().StringVarP(&packageProject, "project", "p", defaultProject, "项目名称（优先使用，如果已在内置列表中，无需指定pom/maven/output）")
	packageCmd.Flags().StringVarP(&packagePom, "pom", "P", "", "pom.xml 路径（当项目未被记录时需手动指定）")
	packageCmd.Flags().StringVarP(&packageMaven, "maven", "m", "", "maven settings.xml 路径（当项目未被记录时需手动指定）")
	packageCmd.Flags().StringVarP(&packageOutput, "output", "o", "", "jar 输出目录（当项目未被记录时需手动指定）")
	packageCmd.Flags().BoolVarP(&packageListProject, "list project", "l", false, "输出内置项目信息")

	//2.返回
	return packageCmd
}

// 打包命令 package java project
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "package java project",
	Run: func(cmd *cobra.Command, args []string) {

		//1.如果打印项目列表，则打印并返回，否则执行打包命令
		if packageListProject {
			consolePackageProjectInfos()
			return
		} else {
			runPackageCommand()
		}
	},
}

// 打印项目信息
func consolePackageProjectInfos() {

	//1.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//2.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 项目名称          \t| 配置类型 | 路径                                                         \t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//3.打印项目信息
	for projectName, projectProperties := range packageCmdProjectPropertiesMap[system] {
		fmt.Fprintf(w, "| %-18s\t| pom    \t| %-60s\t|\n", projectName, projectProperties["pom"])
		fmt.Fprintf(w, "| %-18s\t| maven  \t| %-60s\t|\n", "", projectProperties["maven"])
		fmt.Fprintf(w, "| %-18s\t| output \t| [%-60s]\t|\n", "", projectProperties["output"][1])
		fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	}

	//4.写入控制台
	w.Flush()
}

// 运行打包命令
func runPackageCommand() {

	//1.获取当前系统对应的项目集合
	systemProjects := packageCmdProjectPropertiesMap[system]
	if systemProjects == nil {
		log.Fatalf("当前系统不支持package命令: %v\n", system)
	}

	//2.获取项目配置
	systemProject := systemProjects[packageProject]
	if systemProject == nil {
		log.Fatalf("不支持项目: %v\n", packageProject)
	}

	//3.终止项目进程，确保打包不受影响
	if err := killProjectProcess(systemProject); err != nil {
		log.Fatalf("无法结束项目Java进程: %v\n", err)
	}

	//4.获取maven命令集合
	commands := getMavenCommands(systemProject)

	//5.执行maven命令
	if err := runMavenCommands(commands); err != nil {
		return
	}

	//6.执行完成后，打开目标文件夹
	openOutPutDir(systemProject)
}

// 杀死项目进程
func killProjectProcess(project map[string][]string) error {

	//1.获取kill配置项
	killItems := common.GetProjectItems(project, "kill")

	//2.杀死项目进程
	if err := common.KillProcess(killItems[0], killItems[1]); err != nil {
		return err
	}

	//3.默认返回
	return nil
}

// 获取Maven命令集合
func getMavenCommands(systemProject map[string][]string) [][]string {

	//1.如果自定义项目配置为空，那么使用项目配置覆盖自定义配置
	if packagePom == "" {
		packagePom = common.GetProjectItems(systemProject, "pom")[0]
	}
	if packageMaven == "" {
		packageMaven = common.GetProjectItems(systemProject, "maven")[0]
	}
	if packageOutput == "" {
		packageOutput = common.GetProjectItems(systemProject, "output")[1]
	}

	//2.定义公共参数，减少代码重复
	baseArgs := []string{
		"-f", packagePom,
		"-s", packageMaven,
		"-DskipTests=true",
	}

	//3.按顺序定义三个命令，clean、install、package
	commands := [][]string{
		append([]string{"clean"}, baseArgs...),
		append([]string{"install"}, baseArgs...),
		append([]string{"package"}, baseArgs...),
	}

	//4.返回命令集合
	return commands
}

// 运行Maven命令
func runMavenCommands(commands [][]string) error {

	//1.依次执行命令
	for i, args := range commands {

		//2.打印当前执行的命令序号和参数
		fmt.Printf(">>> 执行第%d条命令: mvn %v\n", i+1, args)

		//3.执行命令，出错则打印并退出
		if err := common.RunCommand("mvn", args...); err != nil {
			fmt.Printf("命令执行失败: %v\n", err)
			return err
		}

		//4.成功时打印提示
		fmt.Printf("第%d条命令执行成功\n\n", i+1)
	}

	//5.所有命令完成的提示
	fmt.Println("所有命令执行完成！")

	//6.默认返回
	return nil
}

// 打开目标文件夹
func openOutPutDir(project map[string][]string) error {

	//1.获取output项
	openCommand := common.GetProjectItems(project, "output")[0]

	//2.执行命令打开文件夹
	if err := common.RunCommand(openCommand, packageOutput); err != nil {
		return err
	}

	//3.默认返回
	return nil
}
