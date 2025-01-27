package main

import (
	"encoding/json"
	"fmt"
	"info/cpuInfo"
	"info/diskInfo"
	"info/memInfo"
	"info/netInfo"
	"info/sysInfo"
	"log"
	"net/http"
	"sync"
)

type DeviceInfo struct {
	CPUFrequencyData CPUMaxMin  `json:"cpuFrequencyData"`
	CPUCurInfoData   CPUCur     `json:"cpuCurInfoData"`
	MemoryData       MemData    `json:"memoryData"`
	DiskData         DiskData   `json:"diskData"`
	NetworkTraffic   NetTraffic `json:"networkTraffic"`
	NetworkInfo      NetInfo    `json:"networkInfo"`
	SysInfoData      SysInfo    `json:"sysInfoData"`
}

type CPUMaxMin struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type CPUCur struct {
	Usage float64 `json:"usage"`
	Temp  float64 `json:"temp"`
	Freq  float64 `json:"freq"`
}

type MemData struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
}

type DiskData struct {
	ModelName        string `json:"modelname"`
	FileSystemFormat string `json:"fsformat"`
	Size             int    `json:"size"`
	Used             int    `json:"used"`
	Available        int    `json:"available"`
	Temp             int    `json:"temp"`
}

type NetTraffic struct {
	Eth0Receive  float64 `json:"eth0receive"`
	Eth0Send     float64 `json:"eth0send"`
	Wlan0Receive float64 `json:"wlan0receive"`
	Wlan0Send    float64 `json:"wlan0send"`
}

type NetInfo struct {
	Hostname   string                     `json:"hostname"`
	Interfaces []netInfo.NetworkInterface `json:"interfaces"`
}

type SysInfo struct {
	Model         string `json:"model"`
	OSVersion     string `json:"osVersion"`
	KernelVersion string `json:"kernelVersion"`
	SystemTime    string `json:"systemTime"`
}

func getDeviceInfo() DeviceInfo {
	var wg sync.WaitGroup

	// 定义变量来存储结果
	var cpuMax, cpuMin float64
	var cpuUsage, cpuTemp, cpuFreq float64
	var memTotal, memAvailable float64
	var diskModelName, fileSystemFormat string
	var diskSize, diskUsed, diskAvail, diskTemp int
	var eth0reci, eth0send, wlan0reci, wlan0send float64
	var hostname string
	var networkinterface []netInfo.NetworkInterface
	var model, osVersion, kernelVersion, systemTime string

	// 启动并发任务
	wg.Add(1)
	go func() {
		defer wg.Done()
		cpuMax, cpuMin = cpuInfo.GetCPUInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cpuUsage, cpuTemp, cpuFreq = cpuInfo.GetCPUCurInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		memTotal, memAvailable = memInfo.GetMemInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		diskModelName, fileSystemFormat, diskSize, diskUsed, diskAvail, diskTemp = diskInfo.GetDiskInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wlan0reci, wlan0send = netInfo.GetNetTraffic("wlan0")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		eth0reci, eth0send = netInfo.GetNetTraffic("eth0")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		hostname, networkinterface = netInfo.GetNetInfo()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		model, osVersion, kernelVersion, systemTime = sysInfo.GetSysInfo()
	}()

	// 等待所有并发任务完成
	wg.Wait()

	return DeviceInfo{
		CPUFrequencyData: CPUMaxMin{
			Max: cpuMax,
			Min: cpuMin,
		},
		CPUCurInfoData: CPUCur{
			Usage: cpuUsage,
			Temp:  cpuTemp,
			Freq:  cpuFreq,
		},
		MemoryData: MemData{
			Total:     memTotal,
			Available: memAvailable,
		},
		DiskData: DiskData{
			ModelName:        diskModelName,
			FileSystemFormat: fileSystemFormat,
			Size:             diskSize,
			Used:             diskUsed,
			Available:        diskAvail,
			Temp:             diskTemp,
		},
		NetworkTraffic: NetTraffic{
			Eth0Receive:  eth0reci,
			Eth0Send:     eth0send,
			Wlan0Receive: wlan0reci,
			Wlan0Send:    wlan0send,
		},
		NetworkInfo: NetInfo{
			Hostname:   hostname,
			Interfaces: networkinterface,
		},
		SysInfoData: SysInfo{
			Model:         model,
			OSVersion:     osVersion,
			KernelVersion: kernelVersion,
			SystemTime:    systemTime,
		},
	}
}

func deviceInfoHandler(w http.ResponseWriter, r *http.Request) {
	deviceInfo := getDeviceInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deviceInfo)
}

func main() {
	// // 下面的信息只打印一次
	// model, osVersion, kernelVersion, systemTime := sysInfo.GetSysInfo()
	// fmt.Printf("Raspberry Pi Model:\t%s\n", model)
	// fmt.Printf("OS Version:\t\t%s\n", osVersion)
	// fmt.Printf("Kernel Version:\t\t%s\n", kernelVersion)
	// fmt.Printf("System Time:\t\t%s\n", systemTime)

	// hostname, netinfoes := netInfo.GetNetInfo()
	// fmt.Println("Hostname: " + hostname)
	// for _, netinfo := range netinfoes {
	// 	fmt.Printf("  Name: %s\t", netinfo.Name)
	// 	fmt.Printf("  IP Address: %s\t", netinfo.IPAddress)
	// 	fmt.Printf("  MAC Address: %s\n", netinfo.MACAddress)
	// }
	// // 下面循环打印
	// for {
	// 	var wg sync.WaitGroup

	// 	// 定义变量来存储结果
	// 	var cpuUsage, cpuTemp, cpuFreq float64
	// 	var memTotal, memAvailable float64
	// 	var diskSize, diskUsed, diskAvail, diskTemp int
	// 	var reci, send float64

	// 	// 启动并发任务
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		cpuUsage, cpuTemp, cpuFreq = cpuInfo.GetCPUInfo()
	// 	}()

	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		memTotal, memAvailable = memInfo.GetMemInfo()
	// 	}()

	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		diskSize, diskUsed, diskAvail, diskTemp = diskInfo.GetDiskInfo()
	// 	}()

	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		reci, send = netInfo.GetNetTraffic("wlan0")
	// 	}()

	// 	// 等待所有并发任务完成
	// 	wg.Wait()

	// 	// 打印结果
	// 	fmt.Printf("CPU Usage:\t%.2f%%,\t\t", cpuUsage)
	// 	fmt.Printf("CPU Temp:\t%.2f C,\t", cpuTemp)
	// 	fmt.Printf("CPU Freq:\t%.2f GHz\n", cpuFreq)

	// 	fmt.Printf("MemTotal:\t%.2f GB,\t", memTotal)
	// 	fmt.Printf("MemAvailable:\t%.2f GB\n", memAvailable)

	// 	fmt.Printf("Disk Total:\t%d GB,\t\t", diskSize)
	// 	fmt.Printf("Disk Used:\t%d GB,\t\t", diskUsed)
	// 	fmt.Printf("Disk Avail:\t%d GB,\t\t", diskAvail)
	// 	fmt.Printf("Temperature:\t%d°C\n", diskTemp)

	// 	fmt.Printf("Receive Bytes:\t%.2f MB/s\tTransmit Bytes:\t%.2f MB/s\n", reci, send)
	// }

	// 提供静态文件服务，包括index.html
	// 提供静态文件服务
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	// 提供设备信息API
	http.HandleFunc("/api/device-info", deviceInfoHandler)

	fmt.Println("Server started at http://<Raspberry-IP>:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
