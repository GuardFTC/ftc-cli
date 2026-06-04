// Package ai @Author:冯铁城 [17615007230@163.com] 2026-06-04 10:00:00
package ai

import "strings"

// ANSI颜色码
const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

// cleanMarkdown 清洗流式chunk中的Markdown语法符号并添加ANSI颜色
func cleanMarkdown(chunk string) string {

	//1.去掉粗体标记 **
	chunk = strings.ReplaceAll(chunk, "**", "")

	//2.处理标题和其他行级格式
	if strings.Contains(chunk, "\n") || strings.HasPrefix(chunk, "#") || strings.HasPrefix(chunk, ">") || strings.HasPrefix(chunk, "【") {
		lines := strings.Split(chunk, "\n")
		for i, line := range lines {
			trimmed := strings.TrimLeft(line, " ")

			//3.Markdown标题 → 黄色加粗【标题】
			if strings.HasPrefix(trimmed, "#### ") {
				lines[i] = colorYellow + colorBold + "【" + strings.TrimPrefix(trimmed, "#### ") + "】" + colorReset
			} else if strings.HasPrefix(trimmed, "### ") {
				lines[i] = colorYellow + colorBold + "【" + strings.TrimPrefix(trimmed, "### ") + "】" + colorReset
			} else if strings.HasPrefix(trimmed, "## ") {
				lines[i] = colorYellow + colorBold + "【" + strings.TrimPrefix(trimmed, "## ") + "】" + colorReset
			} else if strings.HasPrefix(trimmed, "# ") {
				lines[i] = colorYellow + colorBold + "【" + strings.TrimPrefix(trimmed, "# ") + "】" + colorReset
			} else if strings.HasPrefix(trimmed, "【") && strings.HasSuffix(strings.TrimSpace(trimmed), "】") {
				//4.已经是【标题】格式 → 黄色加粗
				lines[i] = colorYellow + colorBold + trimmed + colorReset
			} else if strings.HasPrefix(trimmed, "> ") {
				//5.引用行 → 灰色
				lines[i] = colorGray + trimmed + colorReset
			} else if strings.HasPrefix(trimmed, "```") {
				//6.代码围栏 → 去掉
				lines[i] = ""
			}
		}
		chunk = strings.Join(lines, "\n")
	}

	//7.处理行内代码反引号 → 青色
	if strings.Contains(chunk, "`") {
		result := strings.Builder{}
		inCode := false
		for i := 0; i < len(chunk); i++ {
			if chunk[i] == '`' {
				if inCode {
					result.WriteString(colorReset)
					inCode = false
				} else {
					result.WriteString(colorCyan)
					inCode = true
				}
			} else {
				result.WriteByte(chunk[i])
			}
		}
		if inCode {
			result.WriteString(colorReset)
		}
		chunk = result.String()
	}

	//8.返回
	return chunk
}
