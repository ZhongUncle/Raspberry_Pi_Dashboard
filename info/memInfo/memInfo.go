package memInfo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 定义一个结构体来存储内存信息
type MemoryInfo struct {
	MemTotal     int64
	MemFree      int64
	MemAvailable int64
	Buffers      int64
	Cached       int64
	SwapTotal    int64
	SwapFree     int64
	Active       int64
	Inactive     int64
	AnonPages    int64
	Mapped       int64
	Shmem        int64
	KernelStack  int64
	PageTables   int64
	CommitLimit  int64
	CommittedAS  int64
	VmallocTotal int64
	VmallocUsed  int64
	VmallocChunk int64
}

// 读取 /proc/meminfo 文件并解析内存信息
func readMemInfo(filePath string) (*MemoryInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	memInfo := &MemoryInfo{}

	// 使用正则表达式匹配键值对
	re := regexp.MustCompile(`(\w+):\s+(\d+)\s+kB`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			key := strings.TrimSpace(matches[1])
			value, _ := strconv.ParseInt(matches[2], 10, 64)

			switch key {
			case "MemTotal":
				memInfo.MemTotal = value
			case "MemFree":
				memInfo.MemFree = value
			case "MemAvailable":
				memInfo.MemAvailable = value
			case "Buffers":
				memInfo.Buffers = value
			case "Cached":
				memInfo.Cached = value
			case "SwapTotal":
				memInfo.SwapTotal = value
			case "SwapFree":
				memInfo.SwapFree = value
			case "Active":
				memInfo.Active = value
			case "Inactive":
				memInfo.Inactive = value
			case "AnonPages":
				memInfo.AnonPages = value
			case "Mapped":
				memInfo.Mapped = value
			case "Shmem":
				memInfo.Shmem = value
			case "KernelStack":
				memInfo.KernelStack = value
			case "PageTables":
				memInfo.PageTables = value
			case "CommitLimit":
				memInfo.CommitLimit = value
			case "Committed_AS":
				memInfo.CommittedAS = value
			case "VmallocTotal":
				memInfo.VmallocTotal = value
			case "VmallocUsed":
				memInfo.VmallocUsed = value
			case "VmallocChunk":
				memInfo.VmallocChunk = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return memInfo, nil
}

// 获取内存信息，返回值为（总量， 可使用）
func GetMemInfo() (float64, float64) {
	memInfo, err := readMemInfo("/proc/meminfo")
	if err != nil {
		fmt.Println("Error reading memory info:", err)
	}

	return float64(memInfo.MemTotal), float64(memInfo.MemAvailable)
}
