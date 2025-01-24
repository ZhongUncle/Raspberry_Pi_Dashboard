// // package main

// // import (
// // 	"fmt"
// // 	"info/cpuInfo"
// // 	"info/diskInfo"
// // 	"info/memInfo"
// // )

// // func main() {
// // 	for {
// // 		// 打印CPU信息
// // 		cpuUsage, cpuTmep, cpuFreq := cpuInfo.GetCPUInfo()
// // 		fmt.Printf("CPU Usage:\t%.2f%%,\t\t", cpuUsage)
// // 		fmt.Printf("CPU Temp:\t%.2f C,\t", cpuTmep)
// // 		fmt.Printf("CPU Freq:\t%.2f GHz\n", cpuFreq)
// // 		// 打印内存信息
// // 		memTotal, memAvailable := memInfo.GetMemInfo()
// // 		fmt.Printf("MemTotal:\t%.2f GB,\t", memTotal)
// // 		fmt.Printf("MemAvailable:\t%.2f GB\n", memAvailable)
// // 		// 打印硬盘信息
// // 		diskSize, diskUsed, diskAvail, diskTemp := diskInfo.GetDiskInfo()
// // 		fmt.Printf("Disk Total:\t%d GB,\t\t", diskSize)
// // 		fmt.Printf("Disk Used:\t%d GB,\t\t", diskUsed)
// // 		fmt.Printf("Disk Avail:\t%d GB,\t\t", diskAvail)
// // 		fmt.Printf("Temperature:\t%d°C\n", diskTemp)
// // 	}
// // }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // DiskStats 结构体用于存储磁盘统计信息
// type DiskStats struct {
// 	readBytes  uint64
// 	writeBytes uint64
// }

// // 获取指定磁盘的 I/O 统计信息
// func getDiskStats(device string) (DiskStats, error) {
// 	file, err := os.Open("/proc/diskstats")
// 	if err != nil {
// 		return DiskStats{}, fmt.Errorf("failed to open /proc/diskstats: %v", err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fields := strings.Fields(line)

// 		// 检查是否为目标设备
// 		if len(fields) >= 14 && fields[2] == device {
// 			readBytes, _ := strconv.ParseUint(fields[5], 10, 64) // 读取扇区数
// 			writeBytes, _ := strconv.ParseUint(fields[9], 10, 64) // 写入扇区数

// 			// 扇区大小通常为 512 字节
// 			readBytes *= 512
// 			writeBytes *= 512

// 			return DiskStats{
// 				readBytes:  readBytes,
// 				writeBytes: writeBytes,
// 			}, nil
// 		}
// 	}

// 	return DiskStats{}, fmt.Errorf("device %s not found in /proc/diskstats", device)
// }

// // 计算指定时间间隔内的 I/O 百分比
// func calculateDiskIOPercentage(device string, interval time.Duration) (float64, float64, error) {
// 	// 获取初始统计信息
// 	initialStats, err := getDiskStats(device)
// 	if err != nil {
// 		return 0, 0, fmt.Errorf("failed to get initial stats: %v", err)
// 	}

// 	// 等待指定时间间隔
// 	time.Sleep(interval)

// 	// 获取最终统计信息
// 	finalStats, err := getDiskStats(device)
// 	if err != nil {
// 		return 0, 0, fmt.Errorf("failed to get final stats: %v", err)
// 	}

// 	// 计算 I/O 变化量
// 	readBytesDelta := finalStats.readBytes - initialStats.readBytes
// 	writeBytesDelta := finalStats.writeBytes - initialStats.writeBytes

// 	// 计算 I/O 百分比
// 	totalTime := float64(interval.Seconds())
// 	readPercentage := (float64(readBytesDelta) / (512 * 1024 * 1024)) / totalTime // 以 MB/s 为单位
// 	writePercentage := (float64(writeBytesDelta) / (512 * 1024 * 1024)) / totalTime // 以 MB/s 为单位

// 	return readPercentage, writePercentage, nil
// }

// func main() {
// 	device := "sda" // 替换为目标磁盘设备，例如 "sda" 或 "nvme0n1"
// 	interval := 5 * time.Second

// 	readPercentage, writePercentage, err := calculateDiskIOPercentage(device, interval)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

//		fmt.Printf("Disk %s I/O Percentage:\n", device)
//		fmt.Printf("Read Percentage: %.2f MB/s\n", readPercentage)
//		fmt.Printf("Write Percentage: %.2f MB/s\n", writePercentage)
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// 打开 /proc/net/dev 文件
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个正则表达式，用于匹配网络接口的行
	re := regexp.MustCompile(`\s*(\w+):\s+(\d+)\s+\d+\s+\d+\s+\d+\s+\d+\s+\d+\s+\d+\s+(\d+)\s+`)

	// 使用 bufio.Scanner 逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 使用正则表达式匹配行
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			// 提取接口名称、接收字节数和发送字节数
			interfaceName := matches[1]
			receiveBytes := matches[2]
			transmitBytes := matches[3]

			// 输出结果
			fmt.Printf("Interface: %s\n", interfaceName)
			fmt.Printf("  Receive Bytes: %s\n", receiveBytes)
			fmt.Printf("  Transmit Bytes: %s\n", transmitBytes)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
