package sysInfo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// 获取树莓派型号
func GetRaspberryPiModel() (string, error) {
	file, err := os.Open("/proc/device-tree/model")
	if err != nil {
		return "", fmt.Errorf("failed to open /proc/device-tree/model: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	return "", fmt.Errorf("failed to read model information")
}

// 获取系统版本
func GetOSVersion() (string, error) {
	cmd := exec.Command("cat", "/etc/os-release")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\""), nil
		}
	}

	return "", fmt.Errorf("failed to find PRETTY_NAME in /etc/os-release")
}

// 获取内核版本
func GetKernelVersion() (string, error) {
	cmd := exec.Command("uname", "-r")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// 获取系统信息，返回值为（树莓派型号，获取系统版本，获取内核版本，获取系统时间）
func GetSysInfo() (string, string, string, string) {
	// 获取树莓派型号
	model, err := GetRaspberryPiModel()
	if err != nil {
		fmt.Println("Error getting Raspberry Pi model:", err)
	}

	// 获取系统版本
	osVersion, err := GetOSVersion()
	if err != nil {
		fmt.Println("Error getting OS version:", err)
	}

	// 获取内核版本
	kernelVersion, err := GetKernelVersion()
	if err != nil {
		fmt.Println("Error getting kernel version:", err)
	}

	// 获取系统时间
	systemTime := time.Now().Format("2006-01-02 15:04:05")

	return model, osVersion, kernelVersion, systemTime
}
