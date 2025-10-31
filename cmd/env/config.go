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

// flag变量
var (
	envProject     string
	envListProject bool
)

// env命令 系统-项目名称-项目配置-Map
var envCmdProjectPropertiesMap = map[string]map[string]map[string][]string{
	windows: {
		defaultProject: {
			"nacos": {
				"cmd",
				"/C", "C:\\Users\\Administrator\\base\\nacos\\bin\\startup.cmd",
				"-m", "standalone",
			},
			"sentinel": {
				"java",
				"-Dserver.port=8849",
				"-Dcsp.sentinel.dashboard.server=0.0.0.0:8849",
				"-Dproject.name=Platform",
				"-Dsentinel.dashboard.auth.username=platform",
				"-Dsentinel.dashboard.auth.password=VI7O8ezi18kaYiQupoT2tohAw4mOLi",
				"-jar", "C:\\Users\\Administrator\\base\\sentinel\\sentinel-dashboard-1.8.8.jar",
			},
			"redis": {
				"cmd",
				"/C", `C:\Users\Administrator\base\redis\redis-server.exe`,
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
