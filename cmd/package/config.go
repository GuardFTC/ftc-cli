// Package _package @Author:冯铁城 [17615007230@163.com] 2025-10-31 19:34:04
package _package

import "runtime"

// 定义系统常量
const windows = "windows"
const mac = "darwin"

// 系统名称
var system = runtime.GOOS

// 默认项目
var defaultProject = "prospect-platform"

// package命令 系统-项目名称-项目配置-Map
var packageCmdProjectPropertiesMap = map[string]map[string]map[string][]string{
	windows: {
		defaultProject: {
			"pom":    {"C:\\Users\\Administrator\\project\\java\\prospect-platform\\parent\\pom.xml"},
			"maven":  {"C:\\Users\\Administrator\\maven\\apache-maven-3.9.9-bin\\apache-maven-3.9.9\\conf\\settings.xml"},
			"output": {"explorer", "C:\\Users\\Administrator\\project\\java\\prospect-platform\\output"},
			"kill":   {"java", "prospect."},
		},
		"logging-mon": {
			"pom":    {"C:\\Users\\Administrator\\project\\java\\logging-mon\\pom.xml"},
			"maven":  {"C:\\Users\\Administrator\\maven\\apache-maven-3.9.9-bin\\apache-maven-3.9.9\\conf\\settings.xml"},
			"output": {"explorer", "C:\\Users\\Administrator\\project\\java\\logging-mon\\output"},
			"kill":   {"java", "logging-mon"},
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
