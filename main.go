package main

import (
	"fmt"
	"info/cpuInfo"
)

func main() {
	for {
		fmt.Printf("CPU Usage: %.2f %%\n", cpuInfo.GetCPUUsage())
		fmt.Printf("CPU Temp: %.2f 'C\n", cpuInfo.GetCPUTemp())
	}
}
