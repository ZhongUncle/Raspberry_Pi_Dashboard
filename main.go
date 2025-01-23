package main

import (
	"fmt"
	"info/cpuInfo"
)

func main() {
	for {
		cpuUsage, cpuTmep, cpuFreq := cpuInfo.GetCPUInfo()
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
		fmt.Printf("CPU Temp: %.2f\n", cpuTmep)
		fmt.Printf("CPU Freq: %.2f\n", cpuFreq)
	}
}
