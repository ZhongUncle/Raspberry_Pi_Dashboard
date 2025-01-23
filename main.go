package main

import (
	"fmt"
	"info/cpuInfo"
	"info/memInfo"
)

func main() {
	for {
		cpuUsage, cpuTmep, cpuFreq := cpuInfo.GetCPUInfo()
		fmt.Printf("CPU Usage:\t%.2f%%,\t\t", cpuUsage)
		fmt.Printf("CPU Temp:\t%.2f C,\t", cpuTmep)
		fmt.Printf("CPU Freq:\t%.2f GHz\n", cpuFreq)
		// 打印内存信息
		memTotal, memAvailable := memInfo.GetMemInfo()
		fmt.Printf("MemTotal:\t%.2f GB,\t", memTotal)
		fmt.Printf("MemAvailable:\t%.2f GB\n", memAvailable)
	}
}
