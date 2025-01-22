package main

import (
	"info/cpuInfo"
)

func main() {
	for {
		cpuInfo.GetCPUUsage()
	}
}
