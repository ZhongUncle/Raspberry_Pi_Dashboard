package netInfo

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// 定义一个结构体来存储每个接口的收发字节数
type InterfaceStats struct {
	Name          string
	ReceiveBytes  uint64
	TransmitBytes uint64
}

// 定义一个函数来解析 /proc/net/dev 文件
func ParseNetDev() ([]InterfaceStats, error) {
	// 打开 /proc/net/dev 文件
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// 创建一个缓冲读取器
	scanner := bufio.NewScanner(file)

	// 跳过前两行（标题行）
	for i := 0; i < 2; i++ {
		scanner.Scan()
	}

	// 用于存储所有接口的统计信息
	var stats []InterfaceStats

	// 逐行解析数据
	for scanner.Scan() {
		line := scanner.Text()
		// 按空白字符分割行内容
		fields := strings.Fields(line)

		// 检查字段数量是否符合预期
		if len(fields) < 17 {
			fmt.Println("Unexpected line format:", line)
			continue
		}

		// 提取接口名称（去掉末尾的冒号）
		interfaceName := strings.TrimSuffix(fields[0], ":")

		// 提取接收字节数（第1个字段）
		receiveBytes, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing receive bytes for %s: %v\n", interfaceName, err)
			continue
		}

		// 提取发送字节数（第9个字段）
		transmitBytes, err := strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing transmit bytes for %s: %v\n", interfaceName, err)
			continue
		}

		// 创建结构体实例并存储到切片中
		stats = append(stats, InterfaceStats{
			Name:          interfaceName,
			ReceiveBytes:  receiveBytes,
			TransmitBytes: transmitBytes,
		})
	}

	return stats, nil
}

// 获取所有接口的信息
func PrintAllSocketsInfo() []InterfaceStats {
	// 调用解析函数
	stats, err := ParseNetDev()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return stats
}

// 获取某个接口的信息
func GetSomeSocketInfo(name string) InterfaceStats {
	stats, err := ParseNetDev()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 查找特定接口的统计信息
	for _, stat := range stats {
		if stat.Name == name {
			return stat
		}
	}
	fmt.Println("Wrong Name! Return lo")
	return stats[0]
}

// 获取当前网络收发情况，返回值为收、发速度
func GetNetTraffic(name string) (float64, float64) {
	start := GetSomeSocketInfo(name)
	time.Sleep(1 * time.Second)
	end := GetSomeSocketInfo(name)

	return float64(end.ReceiveBytes-start.ReceiveBytes), float64(end.TransmitBytes-start.TransmitBytes)
}

// 定义结构体存储网络接口信息
type NetworkInterface struct {
	Name       string
	IPAddress  string
	MACAddress string
}

// 获取网络接口信息
func getNetworkInterfaces() ([]NetworkInterface, error) {
	var interfaces []NetworkInterface

	// 获取所有网络接口
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %v", err)
	}

	for _, iface := range ifaces {
		// 跳过非活动接口或虚拟接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的IP地址
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, fmt.Errorf("failed to get addresses for interface %s: %v", iface.Name, err)
		}

		// 获取接口的第一个IPv4地址
		var ipAddress string
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ipAddress = ipNet.IP.String()
					break
				}
			}
		}

		// 获取接口的MAC地址
		macAddress := iface.HardwareAddr.String()

		// 创建结构体实例并存储到切片中
		interfaces = append(interfaces, NetworkInterface{
			Name:       iface.Name,
			IPAddress:  ipAddress,
			MACAddress: macAddress,
		})
	}

	return interfaces, nil
}

// 获得网络信息，返回值为Hostname和接口、IP、MAC。
func GetNetInfo() (string, []NetworkInterface) {
	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
	}

	// 获取所有网络接口信息
	interfaces, err := getNetworkInterfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
	}

	return hostname, interfaces
}
