package cpuInfo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// CPUStat 用于存储CPU的统计信息
type CPUStat struct {
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	Irq       uint64
	SoftIrq   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

// ReadCPUStat 从 /proc/stat 文件中读取CPU统计信息
func ReadCPUStat() (CPUStat, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return CPUStat{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu") {
			parts := strings.Fields(line)
			if len(parts) < 10 {
				continue
			}
			var stat CPUStat
			for i, value := range parts[1:] {
				num, err := strconv.ParseUint(value, 10, 64)
				if err != nil {
					return CPUStat{}, err
				}
				switch i {
				case 0:
					stat.User = num
				case 1:
					stat.Nice = num
				case 2:
					stat.System = num
				case 3:
					stat.Idle = num
				case 4:
					stat.IOWait = num
				case 5:
					stat.Irq = num
				case 6:
					stat.SoftIrq = num
				case 7:
					stat.Steal = num
				case 8:
					stat.Guest = num
				case 9:
					stat.GuestNice = num
				}
			}
			return stat, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return CPUStat{}, err
	}
	return CPUStat{}, nil
}

// CalculateCPUUsage 计算CPU使用率
func CalculateCPUUsage(prevStat, currStat CPUStat) float64 {
	totalPrev := prevStat.User + prevStat.Nice + prevStat.System + prevStat.Idle +
		prevStat.IOWait + prevStat.Irq + prevStat.SoftIrq + prevStat.Steal +
		prevStat.Guest + prevStat.GuestNice

	totalCurr := currStat.User + currStat.Nice + currStat.System + currStat.Idle +
		currStat.IOWait + currStat.Irq + currStat.SoftIrq + currStat.Steal +
		currStat.Guest + currStat.GuestNice

	idlePrev := prevStat.Idle + prevStat.IOWait
	idleCurr := currStat.Idle + currStat.IOWait

	totalDiff := totalCurr - totalPrev
	idleDiff := idleCurr - idlePrev

	if totalDiff == 0 {
		return 0.0
	}

	return 100.0 * (float64(totalDiff-idleDiff) / float64(totalDiff))
}

// 外部使用的函数首字母必须大写，不然认为不能公开使用
// 获取当前CPU利用率
func GetCPUUsage() float64 {
	// 第一次读取CPU统计信息
	prevStat, err := ReadCPUStat()
	if err != nil {
		fmt.Printf("Error reading CPU stats: %v\n", err)
		return -1.0
	}

	// 等待一段时间（例如1秒）
	time.Sleep(1 * time.Second)

	// 第二次读取CPU统计信息
	currStat, err := ReadCPUStat()
	if err != nil {
		fmt.Printf("Error reading CPU stats: %v\n", err)
		return -1.0
	}

	// 计算CPU使用率
	cpuUsage := CalculateCPUUsage(prevStat, currStat)
	return cpuUsage
}

// 获取当前 CPU 温度
func GetCPUTemp() float64 {
	cmd := exec.Command("vcgencmd", "measure_temp")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return -1.0
	}
	tempStr := string(stdout)
	temp, _ := strconv.ParseFloat(tempStr[5:9], 64)
	return temp
}
