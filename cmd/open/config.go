// Package open @Author:冯铁城 [17615007230@163.com] 2026-06-08 17:00:00
package open

import "runtime"

// 定义系统常量
const windows = "windows"
const mac = "darwin"

// 系统名称
var system = runtime.GOOS

// AppConfig 软件配置
type AppConfig struct {
	Command string
	Args    []string
	Always  bool
}

// open命令 系统-软件名称-软件配置-Map
var openCmdAppsMap = map[string]map[string]AppConfig{
	windows: {
		"edge": {
			Command: "cmd",
			Args:    []string{"/c", "start", "msedge", "--restore-last-session"},
			Always:  true,
		},
		"chrome": {
			Command: "cmd",
			Args:    []string{"/c", "start", "chrome", "--restore-last-session"},
			Always:  false,
		},
		"v2ray": {
			Command: "powershell",
			Args:    []string{"-Command", "Start-Process 'D:\\soft\\v2ray\\v2rayN-windows-64\\v2rayN-windows-64\\v2rayN.exe'"},
			Always:  true,
		},
		"docker": {
			Command: "powershell",
			Args:    []string{"-Command", "Start-Process 'C:\\Program Files\\Docker\\Docker\\Docker Desktop.exe'"},
			Always:  true,
		},
		"wechat": {
			Command: "C:\\Program Files\\Tencent\\Weixin\\Weixin.exe",
			Args:    []string{},
			Always:  true,
		},
		"we": {
			Command: "cmd",
			Args:    []string{"/c", "start", "weixin://"},
			Always:  false,
		},
		"idea": {
			Command: "C:\\Program Files\\JetBrains\\IntelliJ IDEA 2025.3\\bin\\idea64.exe",
			Args:    []string{},
			Always:  true,
		},
		"goland": {
			Command: "C:\\Program Files\\JetBrains\\GoLand 2025.2.2\\bin\\goland64.exe",
			Args:    []string{},
			Always:  false,
		},
		"webstorm": {
			Command: "C:\\Program Files\\JetBrains\\WebStorm 2025.2.2\\bin\\webstorm64.exe",
			Args:    []string{},
			Always:  false,
		},
		"datagrip": {
			Command: "C:\\Program Files\\JetBrains\\DataGrip 2025.2.2\\bin\\datagrip64.exe",
			Args:    []string{},
			Always:  true,
		},
		"kiro": {
			Command: "powershell",
			Args:    []string{"-Command", "Start-Process 'kiro'"},
			Always:  true,
		},
		"cursor": {
			Command: "cursor",
			Args:    []string{"."},
			Always:  false,
		},
		"apipost": {
			Command: "powershell",
			Args:    []string{"-Command", "Start-Process 'C:\\Users\\Administrator\\AppData\\Local\\Programs\\Apipost\\Apipost.exe'"},
			Always:  true,
		},
		"yuque": {
			Command: "powershell",
			Args:    []string{"-Command", "Start-Process 'C:\\Users\\Administrator\\AppData\\Local\\Programs\\yuque-desktop\\语雀.exe'"},
			Always:  true,
		},
		"typora": {
			Command: "powershell",
			Args:    []string{"-Command", "Start-Process 'D:\\soft\\typora\\Typora\\Typora.exe'"},
			Always:  true,
		},
		"sublime": {
			Command: "D:\\soft\\sublime\\Sublime Text\\sublime_text.exe",
			Args:    []string{},
			Always:  true,
		},
		"virtual": {
			Command: "C:\\Program Files\\Oracle\\VirtualBox\\VirtualBox.exe",
			Args:    []string{},
			Always:  false,
		},
		"RDM": {
			Command: "C:\\Program Files\\Tiny Craft\\Tiny RDM\\tinyrdm.exe",
			Args:    []string{},
			Always:  false,
		},
		"draw.io": {
			Command: "D:\\soft\\drawio\\draw.io\\draw.io.exe",
			Args:    []string{},
			Always:  false,
		},
	},
	mac: {
		"v2ray": {
			Command: "open",
			Args:    []string{"-a", "v2rayN"},
			Always:  true,
		},
		"edge": {
			Command: "/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
			Args:    []string{"--restore-last-session"},
			Always:  true,
		},
		"wechat": {
			Command: "open",
			Args:    []string{"-a", "WeChat"},
			Always:  true,
		},
		"we": {
			Command: "open",
			Args:    []string{"-a", "WeChat"},
			Always:  true,
		},
		"idea": {
			Command: "open",
			Args:    []string{"-a", "IntelliJ IDEA"},
			Always:  true,
		},
		"goland": {
			Command: "open",
			Args:    []string{"-a", "GoLand"},
			Always:  true,
		},
		"webstorm": {
			Command: "open",
			Args:    []string{"-a", "WebStorm"},
			Always:  true,
		},
		"datagrip": {
			Command: "open",
			Args:    []string{"-a", "DataGrip"},
			Always:  true,
		},
		"kiro": {
			Command: "open",
			Args:    []string{"-a", "kiro"},
			Always:  true,
		},
		"apipost": {
			Command: "open",
			Args:    []string{"-a", "Apipost7"},
			Always:  true,
		},
		"yuque": {
			Command: "open",
			Args:    []string{"-a", "语雀"},
			Always:  true,
		},
		"typora": {
			Command: "open",
			Args:    []string{"-a", "Typora"},
			Always:  true,
		},
		"sublime": {
			Command: "open",
			Args:    []string{"-a", "Sublime Text"},
			Always:  true,
		},
		"docker": {
			Command: "open",
			Args:    []string{"-a", "Docker"},
			Always:  true,
		},
		"draw.io": {
			Command: "open",
			Args:    []string{"-a", "draw.io"},
			Always:  true,
		},
		"ishot": {
			Command: "open",
			Args:    []string{"-a", "iShot"},
			Always:  true,
		},
		"kh": {
			Command: "open",
			Args:    []string{"-a", "KeyboardHolder"},
			Always:  true,
		},
		"arm": {
			Command: "open",
			Args:    []string{"-a", "Another Desktop Manager"},
			Always:  true,
		},
		"tabby": {
			Command: "open",
			Args:    []string{"-a", "Tabby"},
			Always:  false,
		},
	},
}
