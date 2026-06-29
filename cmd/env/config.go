// Package env @Author:冯铁城 [17615007230@163.com] 2025-10-31 19:31:00
package env

import "runtime"

// 定义系统常量
const windows = "windows"
const mac = "darwin"

// 系统名称
var system = runtime.GOOS

// 默认项目
var defaultProject = "prospect-platform"

// env命令 系统-项目名称-项目配置-Map
var envCmdProjectPropertiesMap = map[string]map[string]map[string][]string{
	windows: {
		defaultProject: {
			"nacos": {
				"background",
				"C:\\Users\\Administrator\\project\\java\\logs\\nacos.log",
				"8848",
				"java", "nacos",
				"cmd",
				"/C", "C:\\Users\\Administrator\\base\\nacos\\bin\\startup.cmd",
				"-m", "standalone",
			},
			"sentinel": {
				"background",
				"C:\\Users\\Administrator\\project\\java\\logs\\sentinel.log",
				"8849",
				"java", "sentinel",
				"java",
				"-Dserver.port=8849",
				"-Dcsp.sentinel.dashboard.server=0.0.0.0:8849",
				"-Dproject.name=Platform",
				"-Dsentinel.dashboard.auth.username=platform",
				"-Dsentinel.dashboard.auth.password=VI7O8ezi18kaYiQupoT2tohAw4mOLi",
				"-jar", "C:\\Users\\Administrator\\base\\sentinel\\sentinel-dashboard-1.8.8.jar",
			},
			"redis": {
				"docker",
				"start", "redis-stack",
			},
		},
		"ftcli": {
			"redis": {
				"docker",
				"start", "redis-stack",
			},
			"chroma": {
				"docker",
				"start", "chroma-server",
			},
			"ftcli": {
				"background",
				"C:\\Users\\Administrator\\project\\java\\logs\\ftcli-ai-server.log",
				"6680",
				"java", "ftcli",
				"java",
				"-Dfile.encoding=UTF-8",
				"-Dstdout.encoding=UTF-8",
				"-Dstderr.encoding=UTF-8",
				"-Dhttps.proxyHost=127.0.0.1",
				"-Dhttps.proxyPort=10808",
				"-jar", "C:\\Users\\Administrator\\project\\java\\ftcli-ai-server\\target\\ftcli-ai-server-0.0.1-SNAPSHOT.jar",
			},
		},
		"dolp": {
			"zookeeper": {
				"background",
				"C:\\Users\\Administrator\\project\\java\\logs\\zookeeper.log",
				"2181",
				"java", "zookeeper",
				"cmd",
				"/C", "C:\\Users\\Administrator\\base\\zookeeper\\apache-zookeeper-3.8.6-bin\\bin\\zkServer.cmd",
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
