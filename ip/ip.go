package ip

import (
	"fmt"
	"net"
)

// Ip 获取主机ip的方法
func Ip() string {
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("无法获取网络接口:", err)
		return "127.0.0.1"
	}

	// 遍历每个网络接口
	for _, iface := range interfaces {
		// 排除虚拟和回环接口
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			// 获取接口的地址
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("无法获取接口地址:", err)
				continue
			}
			// 遍历接口的地址
			for _, addr := range addrs {
				// 检查地址是否是 IP 地址
				ipnet, ok := addr.(*net.IPNet)
				if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					// 检查是否是内网地址
					if ipnet.IP.IsGlobalUnicast() && !ipnet.IP.IsLinkLocalUnicast() {
						fmt.Println("内网IP:", ipnet.IP.String())
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	fmt.Println("未找到内网IP")
	return "127.0.0.1"
}
