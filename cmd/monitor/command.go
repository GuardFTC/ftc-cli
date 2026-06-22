// Package monitor @Author:冯铁城 [17615007230@163.com] 2026-06-22 10:00:00
package monitor

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/spf13/cobra"
)

// 常量定义
const (
	lineWidth    = 80
	barMaxBlocks = 24
)

// flag变量
var (
	monitorMemory bool
	monitorCPU    bool
)

// NewMonitorCommand 创建monitor命令
func NewMonitorCommand() *cobra.Command {

	//1.设置Flags
	monitorCmd.Flags().BoolVarP(&monitorMemory, "memory", "m", false, "输出内存信息")
	monitorCmd.Flags().BoolVarP(&monitorCPU, "cpu", "c", false, "输出CPU信息")

	//2.返回
	return monitorCmd
}

// monitorCmd monitor命令
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "system resource monitor (memory & CPU)",
	Run: func(cmd *cobra.Command, args []string) {

		//1.根据flag执行对应操作
		switch {
		case monitorMemory:
			consoleMemoryInfo()
		case monitorCPU:
			consoleCPUInfo()
		default:
			consoleMemoryInfo()
			consoleCPUInfo()
		}
	},
}

// 输出内存信息
func consoleMemoryInfo() {

	//1.获取虚拟内存信息
	vMem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("获取内存信息失败: %v\n", err)
		return
	}

	//2.计算各项内存指标(GB)
	totalGb := float64(vMem.Total) / 1024 / 1024 / 1024
	availableGb := float64(vMem.Available) / 1024 / 1024 / 1024
	usedGb := float64(vMem.Total-vMem.Available) / 1024 / 1024 / 1024
	usedPercent := vMem.UsedPercent

	//3.构建进度条
	bar := buildProgressBar(usedPercent)

	//4.打印标题
	printSectionHeader("MEMORY MONITOR")

	//5.打印内存信息
	fmt.Printf("Total      :  %6.2f GB  [%s]  %.2f%%\n", totalGb, bar, usedPercent)
	fmt.Printf("Used       :  %6.2f GB\n", usedGb)
	fmt.Printf("Available  :  %6.2f GB\n", availableGb)
}

// 输出CPU信息
func consoleCPUInfo() {

	//1.获取物理核心数
	physicalCount, err := cpu.Counts(false)
	if err != nil {
		fmt.Printf("获取物理核心数失败: %v\n", err)
		return
	}

	//2.获取逻辑核心数
	logicalCount, err := cpu.Counts(true)
	if err != nil {
		fmt.Printf("获取逻辑核心数失败: %v\n", err)
		return
	}

	//3.计算系统总CPU利用率(采样3秒)
	totalPercent, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		fmt.Printf("获取系统总CPU利用率失败: %v\n", err)
		return
	}

	//4.计算每个逻辑核心的利用率
	perCpuPercents, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Printf("获取单核CPU利用率失败: %v\n", err)
		return
	}

	//5.打印标题
	printSectionHeader("CPU MONITOR")

	//6.打印核心数与总利用率
	totalUsage := 0.0
	if len(totalPercent) > 0 {
		totalUsage = totalPercent[0]
	}
	fmt.Printf("Cores (P/L): %d / %d%sTotal Usage: %6.2f%%\n",
		physicalCount, logicalCount,
		buildPadding(physicalCount, logicalCount),
		totalUsage)

	//7.按列打印每核心利用率(每行4列)
	printCorePercents(perCpuPercents)

	//8.打印底部分割线
	fmt.Println(strings.Repeat("=", lineWidth))
}

// 构建进度条
func buildProgressBar(percent float64) string {

	//1.计算填充块数
	filled := int(math.Round(percent / 100 * barMaxBlocks))
	if filled > barMaxBlocks {
		filled = barMaxBlocks
	}

	//2.构建进度条字符串
	return strings.Repeat("█", filled) + strings.Repeat("░", barMaxBlocks-filled)
}

// 打印段落标题行
func printSectionHeader(title string) {

	//1.计算标题后需要填充的等号数量
	prefix := "=== " + title + " "
	padLen := lineWidth - len(prefix)
	if padLen < 0 {
		padLen = 0
	}

	//2.打印标题行
	fmt.Printf("%s%s\n", prefix, strings.Repeat("=", padLen))
}

// 构建核心数信息后的填充空格（对齐Total Usage）
func buildPadding(physical, logical int) string {

	//1.计算"Cores (P/L): X / XX"的实际长度
	coresStr := fmt.Sprintf("Cores (P/L): %d / %d", physical, logical)

	//2.计算目标对齐位置（使Total Usage右对齐到行尾）
	// "Total Usage: XXXX.XX%" 长度约为21
	targetPos := lineWidth - 21
	padLen := targetPos - len(coresStr)
	if padLen < 1 {
		padLen = 1
	}

	//3.返回填充空格
	return strings.Repeat(" ", padLen)
}

// 按列打印每核心利用率
func printCorePercents(percents []float64) {

	//1.计算列数和行数(每行4列)
	cols := 4
	total := len(percents)
	rows := (total + cols - 1) / cols

	//2.按行打印
	for row := 0; row < rows; row++ {
		line := ""
		for col := 0; col < cols; col++ {

			//3.计算当前核心索引
			idx := col*rows + row
			if idx >= total {
				break
			}

			//4.格式化核心利用率
			entry := fmt.Sprintf("Core %02d: %6.2f%%", idx, percents[idx])

			//5.拼接列（除最后一列外加分隔符）
			if col > 0 {
				line += " | "
			}
			line += entry
		}

		//6.打印行
		fmt.Println(line)
	}
}
