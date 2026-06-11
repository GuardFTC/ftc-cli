# ftcli

个人开发 CLI 工具，用于自动化 FTC 日常开发工作流：环境启动、项目构建、Maven 打包、CSV 转 SQL、AI 助手、常用软件一键打开。

```
go-ftc-console/
├── main.go              # 入口
├── cmd/
│   ├── root.go          # 根命令 ftcli
│   ├── env/             # 开发环境启动
│   ├── package/         # Java Maven 打包
│   ├── build/           # 构建流水线（kill → package → start）
│   ├── sql/             # CSV → SQL 转换
│   ├── ai/              # AI 助手
│   └── open/            # 常用软件启动
└── common/              # 公共工具（命令执行、进程管理、浏览器打开等）
```

## 环境要求

- Go 1.24+
- Maven 3.x（Java 打包/构建）
- Docker（部分环境服务，如 Redis、Chroma）
- Windows / macOS（配置按系统区分，Windows 配置最完整）

## 安装与构建

```bash
# 克隆项目
git clone <repo-url>
cd go-ftc-console

# 本地运行
go run .

# 编译
go build -o ftcli.exe .

# 或通过 build 命令自举编译（输出到 bin 目录）
ftcli build -p ftcli -t go
```

## 命令总览

| 命令 | 说明 |
|------|------|
| `ftcli env` | 启动项目开发环境（Nacos、Sentinel、Redis 等） |
| `ftcli package` | Java Maven 打包 |
| `ftcli build` | 完整构建流水线：kill → 打包 → 后台启动 |
| `ftcli sql` | 将 CSV 数据转换为 INSERT SQL |
| `ftcli ai` | AI 助手（本地文库 / 网络检索 / 文档上传） |
| `ftcli open` | 一键打开常用开发软件 |

---

## ftcli env — 开发环境

启动项目所需的中间件和后端服务。支持前台命令执行与后台服务启动（含端口检测、进程幂等 kill）。

```bash
ftcli env                          # 启动默认项目（prospect-platform）
ftcli env -p ftcli                 # 启动 ftcli 项目环境
ftcli env -l                       # 列出内置项目及配置
ftcli env -b                       # 查看所有后台服务运行状态
ftcli env --bl ftcli               # 滚动查看指定服务日志
ftcli env --bk ftcli               # 停止指定后台服务
```

**内置项目（Windows）**

| 项目 | 服务 |
|------|------|
| `prospect-platform` | Nacos (8848)、Sentinel (8849)、Redis |
| `ftcli` | Redis、Chroma、ftcli 后端 (6680) |

---

## ftcli package — Java 打包

执行 `mvn clean → install → package`，打包前自动 kill 相关 Java 进程，完成后打开输出目录。

```bash
ftcli package                      # 打包默认项目（prospect-platform）
ftcli package -p ftcli             # 打包 ftcli 后端
ftcli package -p logging-mon       # 打包 logging-mon
ftcli package -l                   # 列出内置项目

# 手动指定路径（项目未在内置列表时）
ftcli package -P <pom路径> -m <settings路径> -o <输出目录>
```

**内置项目（Windows）**：`prospect-platform`、`logging-mon`、`ftcli`

---

## ftcli build — 构建流水线

一键完成 kill → 打包 → 后台启动。支持 Java、Go 及混合构建。

```bash
ftcli build                        # 默认：ftcli 项目，java + go 全部构建
ftcli build -p ftcli -t java       # 仅构建 Java 后端
ftcli build -p ftcli -t go         # 仅编译 Go CLI 自身
ftcli build -p ftcli -t all        # Java + Go 全部构建
ftcli build -l                     # 列出内置项目及支持类型
```

**构建流程**

- **Java**：kill 进程 → Maven 打包 → 后台启动（端口检测）
- **Go**：`go build` 编译；Windows 下通过 bat 脚本延迟替换 exe，避免文件占用

**内置项目（Windows）**：`ftcli`（支持 `java` / `go` / `all`）

---

## ftcli sql — CSV 转 SQL

读取 CSV 文件，按内置表结构生成 `INSERT INTO` 语句。自动处理数字/字符串类型及空值。

```bash
ftcli sql -c data.csv              # 转换 CSV（使用默认库表）
ftcli sql -c data.csv -t <表名> -d <数据库> -o output.sql
ftcli sql -l                       # 列出所有内置表及列名
ftcli sql -p <目录> -c file.csv    # 指定 CSV 所在目录
```

**默认配置**

- 文件路径：`C:\Users\Administrator\Downloads\`
- 默认数据库：`dw_tile`
- 默认表：`ads_bi_af_ltvroas_d_i`
- 默认输出：`output.sql`

内置表涵盖 BI 数据仓库相关表（LTV、留存、素材、付费、A/B 测试等），完整列表见 `cmd/sql/config.go`。

---

## ftcli ai — AI 助手

对接 ftcli 后端服务（默认 `http://localhost:6680`），提供流式聊天、文档上传、管理页面。

```bash
ftcli ai -l                        # 本地文库流式聊天
ftcli ai -w                        # 网络检索流式聊天
ftcli ai -u <路径>                 # 上传指定文件/目录
ftcli ai -f                        # 打开文档管理页面
ftcli ai -t                        # 打开工具管理页面
ftcli ai -s http://localhost:6680  # 指定后端地址
```

聊天模式中输入 `exit` 或按 `Ctrl+C` 退出，输入 `clear` 清屏。

---

## ftcli open — 常用软件

一键启动日常开发软件。默认启动 `always=true` 的软件。

```bash
ftcli open                         # 启动所有默认软件
ftcli open -l                      # 列出支持的软件
ftcli open goland                  # 启动指定软件
ftcli open goland webstorm         # 启动多个软件（空格分隔）
```

**内置软件（Windows）**

| 软件 | 默认启动 |
|------|----------|
| edge、v2ray、docker、wechat、idea、datagrip、kiro、apipost、yuque、typora、sublime | true |
| chrome、goland、webstorm、cursor、wechat_wake | false |

---

## 配置说明

各子命令的项目/软件配置集中在对应 `config.go` 中，按 `runtime.GOOS` 区分系统：

| 模块 | 配置文件 |
|------|----------|
| env | `cmd/env/config.go` |
| package | `cmd/package/config.go` |
| build | `cmd/build/config.go` |
| sql | `cmd/sql/config.go` |
| open | `cmd/open/config.go` |
| ai | `cmd/ai/config.go` |

修改路径、端口、进程关键字等，直接编辑对应 `config.go` 即可，无需改动业务逻辑。

---

## 技术栈

- [Cobra](https://github.com/spf13/cobra) — CLI 框架
- [gopsutil](https://github.com/shirou/gopsutil) — 进程管理
- [glamour](https://github.com/charmbracelet/glamour) — 终端 Markdown 渲染（AI 模块）

## 版本

当前版本：`1.0.0`
