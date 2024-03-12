package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

// RegisterService
// 注册实例
func RegisterService(config *ClientConfig, instanceConfig vo.RegisterInstanceParam) bool {
	var err error
	namingClient, err := ServiceClient(config)
	if err != nil {
		log.Println(err)
		return false
	}
	success, err := namingClient.RegisterInstance(instanceConfig)
	if err != nil {
		log.Println(err)
		return false
	}
	return success
}

// GetService
// 获取实例详情
func GetService(config *ClientConfig, serviceInfo vo.GetServiceParam) (model.Service, error) {
	client, err := ServiceClient(config)
	if err != nil {
		log.Println(err)
		return model.Service{}, err
	}
	service, err := client.GetService(serviceInfo)
	return service, nil
}

// GetAllServices
// 获取所有实例
func GetAllServices(config *ClientConfig, instanceConfig vo.SelectAllInstancesParam) ([]model.Instance, error) {
	client, err := ServiceClient(config)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	services, err := client.SelectAllInstances(instanceConfig)
	return services, nil
}

// ServiceClient
// 创建注册中心客户端
func ServiceClient(config *ClientConfig) (naming_client.INamingClient, error) {
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

	// 创建服务发现客户端
	var namingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Println("创建nacos服务注册中心客户端失败！")
		return nil, err
	}
	return namingClient, nil
}
