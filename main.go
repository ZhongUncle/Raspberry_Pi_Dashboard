package main

import (
	"fmt"
	"info/cpuInfo"
	"info/diskInfo"
	"info/memInfo"
	"info/netInfo"
	"info/sysInfo"
	"sync"
)

func main() {
	// 下面的信息只打印一次
	model, osVersion, kernelVersion, systemTime := sysInfo.GetSysInfo()
	fmt.Printf("Raspberry Pi Model:\t%s\n", model)
	fmt.Printf("OS Version:\t\t%s\n", osVersion)
	fmt.Printf("Kernel Version:\t\t%s\n", kernelVersion)
	fmt.Printf("System Time:\t\t%s\n", systemTime)

	hostname, netinfoes := netInfo.GetNetInfo()
	fmt.Println("Hostname: " + hostname)
	for _, netinfo := range netinfoes {
		fmt.Printf("  Name: %s\t", netinfo.Name)
		fmt.Printf("  IP Address: %s\t", netinfo.IPAddress)
		fmt.Printf("  MAC Address: %s\n", netinfo.MACAddress)
	}
	// 下面循环打印
	for {
		var wg sync.WaitGroup

		// 定义变量来存储结果
		var cpuUsage, cpuTemp, cpuFreq float64
		var memTotal, memAvailable float64
		var diskSize, diskUsed, diskAvail, diskTemp int
		var reci, send float64

		// 启动并发任务
		wg.Add(1)
		go func() {
			defer wg.Done()
			cpuUsage, cpuTemp, cpuFreq = cpuInfo.GetCPUInfo()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			memTotal, memAvailable = memInfo.GetMemInfo()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			diskSize, diskUsed, diskAvail, diskTemp = diskInfo.GetDiskInfo()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			reci, send = netInfo.GetNetTraffic("wlan0")
		}()

		// 等待所有并发任务完成
		wg.Wait()

		// 打印结果
		fmt.Printf("CPU Usage:\t%.2f%%,\t\t", cpuUsage)
		fmt.Printf("CPU Temp:\t%.2f C,\t", cpuTemp)
		fmt.Printf("CPU Freq:\t%.2f GHz\n", cpuFreq)

		fmt.Printf("MemTotal:\t%.2f GB,\t", memTotal)
		fmt.Printf("MemAvailable:\t%.2f GB\n", memAvailable)

		fmt.Printf("Disk Total:\t%d GB,\t\t", diskSize)
		fmt.Printf("Disk Used:\t%d GB,\t\t", diskUsed)
		fmt.Printf("Disk Avail:\t%d GB,\t\t", diskAvail)
		fmt.Printf("Temperature:\t%d°C\n", diskTemp)

		fmt.Printf("Receive Bytes:\t%.2f MB/s\tTransmit Bytes:\t%.2f MB/s\n", reci, send)
	}
}
