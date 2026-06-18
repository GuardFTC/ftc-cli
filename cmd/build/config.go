// Package build @Author:冯铁城 [17615007230@163.com] 2026-06-08 16:00:00
package build

import "runtime"

// 定义系统常量
const windows = "windows"
const mac = "darwin"

// 系统名称
var system = runtime.GOOS

// 默认项目
var defaultProject = "ftcli"

// 默认构建类型
var defaultType = "all"

// 构建类型常量
const typeJava = "java"
const typeGo = "go"
const typeAll = "all"

// build命令 项目名称-构建类型-配置项-Map
var buildCmdProjectPropertiesMap = map[string]map[string]map[string][]string{
	windows: {
		"ftcli": {
			"java-kill":  {"java", "ftcli"},
			"java-pom":   {"C:\\Users\\Administrator\\project\\java\\ftcli\\pom.xml"},
			"java-maven": {"C:\\Users\\Administrator\\maven\\apache-maven-3.9.9-bin\\apache-maven-3.9.9\\conf\\settings.xml"},
			"java-log":   {"C:\\Users\\Administrator\\project\\java\\logs\\ftcli-server.log"},
			"java-port":  {"6680"},
			"java-start": {
				"java",
				"-Dfile.encoding=UTF-8",
				"-Dstdout.encoding=UTF-8",
				"-Dstderr.encoding=UTF-8",
				"-Dhttps.proxyHost=127.0.0.1",
				"-Dhttps.proxyPort=10808",
				"-jar", "C:\\Users\\Administrator\\project\\java\\ftcli\\target\\ftcli-0.0.1-SNAPSHOT.jar",
			},
			"go-source": {"C:\\Users\\Administrator\\project\\go\\src\\ftcli"},
			"go-output": {"..\\..\\bin\\ftcli.exe"},
		},
	},
	mac: {
		"ftcli": {
			"java-kill":  {"java", "ftcli"},
			"java-pom":   {},
			"java-maven": {},
			"java-start": {},
			"go-source":  {},
			"go-output":  {},
		},
	},
}
