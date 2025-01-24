package main

import (
	"fmt"
	"info/cpuInfo"
	"info/diskInfo"
	"info/memInfo"
	"info/netInfo"
)

func main() {
	hostname, netinfoes := netInfo.GetNetInfo()
	fmt.Println("Hostname: " + hostname)
	for _, netinfo := range netinfoes {
		fmt.Printf("  Name: %s\t", netinfo.Name)
		fmt.Printf("  IP Address: %s\t", netinfo.IPAddress)
		fmt.Printf("  MAC Address: %s\n", netinfo.MACAddress)
	}
	for {
		// 打印CPU信息
		cpuUsage, cpuTmep, cpuFreq := cpuInfo.GetCPUInfo()
		fmt.Printf("CPU Usage:\t%.2f%%,\t\t", cpuUsage)
		fmt.Printf("CPU Temp:\t%.2f C,\t", cpuTmep)
		fmt.Printf("CPU Freq:\t%.2f GHz\n", cpuFreq)
		// 打印内存信息
		memTotal, memAvailable := memInfo.GetMemInfo()
		fmt.Printf("MemTotal:\t%.2f GB,\t", memTotal)
		fmt.Printf("MemAvailable:\t%.2f GB\n", memAvailable)
		// 打印硬盘信息
		diskSize, diskUsed, diskAvail, diskTemp := diskInfo.GetDiskInfo()
		fmt.Printf("Disk Total:\t%d GB,\t\t", diskSize)
		fmt.Printf("Disk Used:\t%d GB,\t\t", diskUsed)
		fmt.Printf("Disk Avail:\t%d GB,\t\t", diskAvail)
		fmt.Printf("Temperature:\t%d°C\n", diskTemp)
		// 打印网络信息
		reci, send := netInfo.GetNetTraffic("wlan0")
		fmt.Printf("Receive Bytes:\t%.2f MB/s\tTransmit Bytes:\t%.2f MB/s\n", reci, send)
	}
}
