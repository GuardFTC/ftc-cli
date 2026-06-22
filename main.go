package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {

	//1.初始化命令行工具
	//cmd.Init()

	// 1. 输出内存信息
	consoleMemoryInfo()

	// 2. 输出 CPU 信息
	consoleCPUInfo()
}

// consoleMemoryInfo 输出内存信息
func consoleMemoryInfo() {
	// 获取虚拟内存信息（包含总内存、可用内存等）
	vMem, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("获取内存信息失败: %v\n", err)
		return
	}

	total := vMem.Total
	available := vMem.Available
	used := total - available

	// 转换为 GB (浮点数计算，避免整除截断)
	totalGb := float64(total) / 1024 / 1024 / 1024
	availableGb := float64(available) / 1024 / 1024 / 1024
	usedGb := float64(used) / 1024 / 1024 / 1024

	fmt.Println("=== Gopsutil 内存监控 ===")
	fmt.Printf("总内存: %.2f GB\n", totalGb)
	fmt.Printf("可用内存: %.2f GB\n", availableGb)
	fmt.Printf("已用内存: %.2f GB\n", usedGb)
	// vMem.UsedPercent 已经由底层帮我们计算好了
	fmt.Printf("内存使用率: %.2f%%\n", vMem.UsedPercent)
}

// consoleCPUInfo 输出 CPU 信息
func consoleCPUInfo() {
	fmt.Println("\n=== Gopsutil CPU 监控 ===")

	// 1. 获取物理核心数与逻辑核心数
	physicalCount, err := cpu.Counts(false)
	if err != nil {
		fmt.Printf("获取物理核心数失败: %v\n", err)
		return
	}
	logicalCount, err := cpu.Counts(true)
	if err != nil {
		fmt.Printf("获取逻辑核心数失败: %v\n", err)
		return
	}

	fmt.Printf("CPU 物理核心数: %d\n", physicalCount)
	fmt.Printf("CPU 逻辑核心数: %d\n", logicalCount)

	// 2. 计算系统总 CPU 利用率
	// cpu.Percent 的第一个参数是采样等待时间，传入 5*time.Second 代表阻塞 5 秒进行采样
	// 第二个参数如果为 false，返回包含一个元素的切片，代表总利用率
	totalPercent, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		fmt.Printf("获取系统总 CPU 利用率失败: %v\n", err)
		return
	}
	if len(totalPercent) > 0 {
		fmt.Printf("系统总 CPU 利用率: %.2f%%\n", totalPercent[0])
	}

	// 3. 计算每个逻辑核心的利用率
	// 第二个参数如果为 true，会返回每个核心独立的利用率切片
	// 注意：因为上面已经阻塞了 5 秒，这里传入 0 即可复用或进行瞬时计算（若想更准也可以同样传入 time.Second）
	perCpuPercents, err := cpu.Percent(0, true)
	if err != nil {
		fmt.Printf("获取单核 CPU 利用率失败: %v\n", err)
		return
	}

	for i, percent := range perCpuPercents {
		fmt.Printf("核心 %d 利用率: %.2f%%\n", i, percent)
	}
}
