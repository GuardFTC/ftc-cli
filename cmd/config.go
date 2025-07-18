// Package cmd @Author:冯铁城 [17615007230@163.com] 2025-07-14 17:01:53
package cmd

import "runtime"

// 定义系统常量
const windows = "windows"
const mac = "darwin"

// 系统名称
var system = runtime.GOOS

// 默认项目
var defaultProject = "prospect-platform"

// flag变量
var (

	// package命令相关flag变量
	packageProject     string
	packagePom         string
	packageMaven       string
	packageOutput      string
	packageListProject bool

	// env命令相关flag变量
	envProject     string
	envListProject bool
)

// package命令 系统-项目名称-项目配置-Map
var packageCmdProjectPropertiesMap = map[string]map[string]map[string][]string{
	windows: {
		defaultProject: {
			"pom":    {"D:/project/java/prospect-platform/parent/pom.xml"},
			"maven":  {"D:/base/maven/apache-maven-3.9.9-bin/apache-maven-3.9.9/conf/settings.xml"},
			"output": {"explorer", "D:\\project\\java\\prospect-platform\\output"},
			"kill":   {"java", "prospect."},
		},
	},
	mac: {
		defaultProject: {
			"pom":    {},
			"maven":  {},
			"output": {"open"},
			"kill":   {"java", "prospect."},
		},
	},
}

// env命令 系统-项目名称-项目配置-Map
var envCmdProjectPropertiesMap = map[string]map[string]map[string][]string{
	windows: {
		defaultProject: {
			"nacos": {
				"cmd",
				"/C", `D:\base\nacos\bin\startup.cmd`,
				"-m", "standalone",
			},
			"sentinel": {
				"java",
				"-Dserver.port=8849",
				"-Dcsp.sentinel.dashboard.server=0.0.0.0:8849",
				"-Dproject.name=Platform",
				"-Dsentinel.dashboard.auth.username=platform",
				"-Dsentinel.dashboard.auth.password=VI7O8ezi18kaYiQupoT2tohAw4mOLi",
				"-jar", "D:\\base\\sentinel\\sentinel-dashboard-1.8.8.jar",
			},
		},
	},
	mac: {
		defaultProject: {
			"nacos":    {},
			"sentinel": {},
		},
	},
}
