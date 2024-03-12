package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"gopkg.in/yaml.v2"
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// GetRemoteConfig 获取远程配置
// 泛型声明配置类结构体类型
// 参数如泛型 T 声明的结构体变量
// 返回值为参数的实例
func GetRemoteConfig(config *ClientConfig, result any) error {
	configClient, err := ConfigClient(config)
	if err != nil {
		log.Println("创建nacos配置客户端失败！")
		return err
	}
	// 获取配置信息
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: *config.DataId,
		Group:  *config.Group,
	})
	if err != nil {
		log.Println("naocs获取动态配置失败！")
		return err
	}
	// yaml解析并赋值
	err = yaml.Unmarshal([]byte(content), result)
	if err != nil {
		log.Println("yaml解析配置失败！")
		return err
	}
	fmt.Printf("获取到的配置信息：%#v\n", result)
	return nil
}

func ConfigClient(config *ClientConfig) (config_client.IConfigClient, error) {
	// 配置nacos clientconfig
	var clientConfig = constant.ClientConfig{
		NamespaceId:         *config.NamespaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           *config.TimeoutMs,
		NotLoadCacheAtStart: *config.NotLoadCacheAtStart,
		LogDir:              *config.LogDir,
		CacheDir:            *config.CacheDir,
		LogLevel:            *config.LogLevel,
		Username:            *config.Conn.Username,
		Password:            *config.Conn.Password,
	}

	// ServerConfig
	var serverConfigs = []constant.ServerConfig{
		{
			IpAddr:      *config.Conn.IpAddr,
			ContextPath: *config.Conn.ContextPath,
			Port:        *config.Conn.Port,
			Scheme:      *config.Conn.Scheme,
		},
	}

	// 创建动态配置客户端
	var configClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Println("创建nacos配置客户端失败！")
		return nil, err
	}
	return configClient, nil
}
