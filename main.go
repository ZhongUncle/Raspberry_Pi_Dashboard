package main

import (
	"fmt"
	"info/cpuInfo"
)

func main() {
	for {
		fmt.Printf("CPU Temp: %s\n", cpuInfo.GetCPUTemp())
	}
}
