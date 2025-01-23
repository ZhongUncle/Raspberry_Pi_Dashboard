package diskInfo

import (
	"bufio"
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// 获取硬盘使用情况（返回值为Size、Used 和 Avail）
func GetDiskUsage() (int, int, int) {
	// 执行 df -h 命令
	cmd := exec.Command("df", "-h")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return -1, -1, -1
	}

	// 解析输出
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		// 跳过标题行
		if strings.HasPrefix(line, "Filesystem") {
			continue
		}

		// 分割行
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		mountPoint := fields[len(fields)-1]
		if mountPoint == "/" {
			// 提取 Size、Used 和 Avail
			sizeStr := fields[1]
			usedStr := fields[2]
			availStr := fields[3]

			size := 0
			for i := 0; i < len(sizeStr)-1; i++ {
				size = size*10 + (int(sizeStr[i]) - '0')
			}
			used := 0
			for i := 0; i < len(usedStr)-1; i++ {
				used = used*10 + (int(usedStr[i]) - '0')
			}
			avail := 0
			for i := 0; i < len(availStr)-1; i++ {
				avail = avail*10 + (int(availStr[i]) - '0')
			}

			return size, used, avail
		}
	}

	return -1, -1, -1
}

// 获取硬盘温度
func GetDiskTemperature() int {
	// 执行 smartctl -a 命令
	cmd := exec.Command("sudo", "smartctl", "-a", "/dev/sda")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0
	}

	// 使用正则表达式匹配目标行
	re := regexp.MustCompile(`^194 Temperature_Celsius\s+0x[0-9a-fA-F]+\s+\d+\s+\d+\s+\d+\s+\w+\s+\w+\s+-\s+(\d+)$`)
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches != nil && len(matches) > 1 {
			// 提取最后一个整数值
			temp, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0
			}
			return temp
		}
	}

	if err := scanner.Err(); err != nil {
		return 0
	}

	return 0
}

func GetDiskInfo() (int, int, int, int) {
	size, used, avail := GetDiskUsage()
	temp := GetDiskTemperature()

	return size, used, avail, temp
}
