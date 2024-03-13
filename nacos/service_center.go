package nacos

import (
	"github.com/flairamos/go-component/xlog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// RegisterInstance
// 注册实例
func RegisterInstance(config *ClientConfig, instanceConfig vo.RegisterInstanceParam) bool {
	var err error
	namingClient, err := GetNamingClient(config)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return false
	}
	success, err := namingClient.RegisterInstance(instanceConfig)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return false
	}
	return success
}

// DeregisterInstance
// 注销实例
func DeregisterInstance(config *ClientConfig, instanceConfig vo.DeregisterInstanceParam) bool {
	var err error
	namingClient, err := GetNamingClient(config)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return false
	}
	instance, err := namingClient.DeregisterInstance(instanceConfig)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return false
	}
	return instance
}

// GetService
// 获取服务信息
func GetService(config *ClientConfig, serviceInfo vo.GetServiceParam) (model.Service, error) {
	client, err := GetNamingClient(config)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return model.Service{}, err
	}
	service, err := client.GetService(serviceInfo)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return model.Service{}, err
	}
	return service, nil
}

// SelectInstances
// 获取所有实例 (可以筛选健康实例)
func SelectInstances(config *ClientConfig, instanceConfig vo.SelectInstancesParam) ([]model.Instance, error) {
	client, err := GetNamingClient(config)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return nil, err
	}
	services, err := client.SelectInstances(instanceConfig)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return nil, err
	}
	return services, nil
}

// SelectOneHealthyInstance
// 获取健康的实例
// health=true,enable=true and weight>0
func SelectOneHealthyInstance(config *ClientConfig, serviceInfo vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	client, err := GetNamingClient(config)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return nil, err
	}
	services, err := client.SelectOneHealthyInstance(serviceInfo)
	if err != nil {
		xlog.GoLogger.ErrorP(err)
		return nil, err
	}
	return services, nil
}

// GetNamingClient
// 创建注册中心客户端
func GetNamingClient(config *ClientConfig) (naming_client.INamingClient, error) {
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
		xlog.GoLogger.ErrorP("创建nacos服务注册中心客户端失败！")
		return nil, err
	}
	return namingClient, nil
}
