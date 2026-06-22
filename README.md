# ftcli

个人开发 CLI 工具，用于自动化日常开发工作流：环境启动、项目构建、Maven 打包、CSV 转 SQL、AI 助手、常用软件一键打开。

```
ftcli/
├── main.go              # 入口
├── cmd/
│   ├── root.go          # 根命令 ftcli (v1.0.0)
│   ├── ai/              # AI 助手（对接 ftcli 后端）
│   ├── build/           # 构建流水线（kill → package → start）
│   ├── env/             # 开发环境启动
│   ├── monitor/         # 系统资源监控（内存/CPU）
│   ├── open/            # 常用软件启动
│   ├── package/         # Java Maven 打包
│   └── sql/             # CSV → SQL 转换
└── util/                # 公共工具（命令执行、进程管理、浏览器打开等）
```

## 环境要求

- Go 1.24+
- Maven 3.x（Java 打包/构建需要）
- Docker（部分环境服务依赖容器）
- Windows 为主要支持平台，macOS 有预留配置结构

## 安装与构建

```bash
# 本地运行
go run .

# 编译为可执行文件
go build -o ftcli.exe .

# 通过 build 命令自举编译（输出到 ../../bin/ftcli.exe）
ftcli build -p ftcli -t go
```

## 命令总览

| 命令 | 说明 |
|------|------|
| `ftcli env` | 启动项目开发环境 |
| `ftcli package` | Java Maven 打包 |
| `ftcli build` | 完整构建流水线：kill → 打包 → 后台启动 |
| `ftcli sql` | CSV 数据转换为 INSERT SQL |
| `ftcli ai` | AI 助手（流式聊天 / 文档上传 / 管理页面） |
| `ftcli open` | 一键打开常用开发软件 |
| `ftcli monitor` | 系统资源监控（内存/CPU） |

---

## ftcli env

启动项目所需的中间件和后端服务。支持 Docker 容器启动与后台进程启动（含端口检测、进程幂等 kill）。

```bash
ftcli env                  # 启动默认项目（prospect-platform）
ftcli env -p ftcli         # 启动 ftcli 项目环境
ftcli env -l               # 列出内置项目及配置
ftcli env -b               # 查看所有后台服务运行状态
ftcli env --bl ftcli       # 滚动查看指定服务日志
ftcli env --bk ftcli       # 停止指定后台服务
```

**内置项目（Windows）**

| 项目 | 启动服务 |
|------|----------|
| `prospect-platform` | Nacos (8848)、Sentinel (8849)、Redis (Docker) |
| `ftcli` | Redis (Docker)、Chroma (Docker)、ftcli 后端 (6680) |

---

## ftcli package

执行 Maven `clean → install → package`，打包前自动 kill 相关 Java 进程，完成后打开输出目录。

```bash
ftcli package              # 打包默认项目（prospect-platform）
ftcli package -p ftcli     # 打包 ftcli 后端
ftcli package -p logging-mon
ftcli package -l           # 列出内置项目

# 手动指定路径（项目未在内置列表时）
ftcli package -P <pom路径> -m <settings路径> -o <输出目录>
```

**内置项目（Windows）**：`prospect-platform`、`logging-mon`、`ftcli`

---

## ftcli build

一键完成 kill → Maven 打包 → 后台启动。支持 Java、Go 及混合构建。

```bash
ftcli build                # 默认：ftcli 项目，java + go 全部构建
ftcli build -p ftcli -t java
ftcli build -p ftcli -t go
ftcli build -p ftcli -t all
ftcli build -l             # 列出内置项目及支持类型
```

**构建流程**

- **Java**：kill 进程 → Maven 打包 → 后台启动（端口存活检测）
- **Go**：编译到临时文件；Windows 下通过 bat 脚本等当前进程退出后延迟替换 exe

**内置项目（Windows）**：`ftcli`（支持 java / go / all）

---

## ftcli sql

读取 CSV 文件，按内置表结构生成 `INSERT INTO` 语句。自动处理数字/字符串类型识别及空值。

```bash
ftcli sql -c data.csv                  # 使用默认库表转换
ftcli sql -c data.csv -t <表名> -d <数据库> -o output.sql
ftcli sql -l                           # 列出所有内置表
ftcli sql -p <目录> -c file.csv        # 指定 CSV 所在目录
```

**默认配置**

| 配置 | 值 |
|------|-----|
| CSV 目录 | `C:\Users\Administrator\Downloads\` |
| 数据库 | `dw_tile` |
| 表 | `ads_bi_af_ltvroas_d_i` |
| 输出 | `output.sql` |

内置 30+ 张 BI 数据仓库表结构（LTV、留存、素材、付费、A/B 测试、活跃、流失等），完整列表见 `cmd/sql/config.go`。

---

## ftcli ai

对接 ftcli Java 后端服务（默认 `http://localhost:6680`），提供流式聊天、文档上传、管理页面入口。

```bash
ftcli ai -l                # 本地文库流式聊天
ftcli ai -w                # 网络检索流式聊天
ftcli ai -u <路径>         # 上传文档到知识库
ftcli ai -f                # 浏览器打开文档管理页面
ftcli ai -t                # 浏览器打开工具管理页面
ftcli ai -s                # 浏览器打开技能管理页面
ftcli ai -S <地址>         # 指定后端地址（默认 localhost:6680）
```

聊天模式：输入 `exit` 或 `Ctrl+C` 退出，`clear` 清屏。输出自动做 Markdown → ANSI 颜色转换（标题黄色、代码青色、引用灰色）。

---

## ftcli open

一键启动日常开发软件。无参数时启动所有 `Always=true` 的软件，也可按名称指定。

```bash
ftcli open                 # 启动所有默认软件
ftcli open -l              # 列出支持的软件及默认启动状态
ftcli open goland          # 启动指定软件（空格分隔多个）
ftcli open goland webstorm
```

**内置软件（Windows）**

| 默认启动 | 软件 |
|----------|------|
| ✓ | edge、v2ray、docker、wechat、idea、datagrip、kiro、apipost、yuque、typora、sublime |
| ✗ | chrome、goland、webstorm、cursor、we、virtual、RDM、draw.io |

---

## ftcli monitor

系统资源监控，输出内存和 CPU 使用情况，带进度条和多核利用率分列展示。

```bash
ftcli monitor              # 输出内存 + CPU 全部信息
ftcli monitor -m           # 仅输出内存信息
ftcli monitor -c           # 仅输出 CPU 信息
```

**输出示例**

```
=== MEMORY MONITOR =============================================================
Total      :   31.77 GB  [████████████████░░░░░░░░]  80.00%
Used       :   25.61 GB
Available  :    6.16 GB
=== CPU MONITOR ================================================================
Cores (P/L): 12 / 20                                       Total Usage:   2.01%
Core 00:   0.52% | Core 05:   0.00% | Core 10:   1.04% | Core 15:   0.00%
Core 01:   0.00% | Core 06:   0.00% | Core 11:   0.00% | Core 16:   5.18%
Core 02:   0.00% | Core 07:   0.00% | Core 12:   0.00% | Core 17:   0.52%
Core 03:   0.00% | Core 08:   4.66% | Core 13:   0.00% | Core 18:  15.03%
Core 04:   0.00% | Core 09:   0.00% | Core 14:   1.04% | Core 19:  12.44%
================================================================================
```

---

## 配置说明

各子命令的配置集中在对应 `config.go` 中，按 `runtime.GOOS` 区分系统。修改路径、端口、进程关键字等直接编辑即可，无需改业务逻辑。

| 模块 | 配置文件 |
|------|----------|
| env | `cmd/env/config.go` |
| package | `cmd/package/config.go` |
| build | `cmd/build/config.go` |
| sql | `cmd/sql/config.go` |
| ai | `cmd/ai/config.go` |
| open | `cmd/open/config.go` |

---

## 依赖

| 库 | 用途 |
|----|------|
| [cobra](https://github.com/spf13/cobra) | CLI 框架 |
| [gopsutil/v3](https://github.com/shirou/gopsutil) | 进程管理（kill、端口检测） |
| [gopsutil/v4](https://github.com/shirou/gopsutil) | 系统资源监控（内存/CPU） |

---

## 作者

冯铁城 — 17615007230@163.com
