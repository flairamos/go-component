package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"testing"
)

type RpcConfig struct {
	Env      string         `yaml:"env"`
	Kitex    KitexLogConfig `yaml:"kitex"`
	Registry RegistryConfig `yaml:"registry"`
	MySQL    MySQLConfig    `yaml:"mysql"`
	Redis    RedisConfig    `yaml:"redis"`
}

type KitexLogConfig struct {
	Service       string `yaml:"service"`
	Address       string `yaml:"address"`
	LogLevel      string `yaml:"log_level"`
	LogFileName   string `yaml:"log_file_name"`
	LogMaxSize    int    `yaml:"log_max_size"`
	LogMaxAge     int    `yaml:"log_max_age"`
	LogMaxBackups int    `yaml:"log_max_backups"`
}

type RegistryConfig struct {
	RegistryAddress []string `yaml:"registry_address"`
	Username        string   `yaml:"username"`
	Password        string   `yaml:"password"`
}

type MySQLConfig struct {
	DSN string `yaml:"dsn"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func TestGetRemoteConfigTest(t *testing.T) {

	instance := DefaultClient("public", "mysql", "dev_food_platform", nil)
	t.Log(*instance.Conn.Username, *instance.NamespaceId, *instance.Group)
	var result RpcConfig
	err := GetRemoteConfig(instance, &result)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestGetRemoteConfig2(t *testing.T) {
	instance := NewRemoteInstance("nacos.nacostest.xyz", 443, "nacos", "yeduGJ4710")
	fmt.Println(instance)
	config := DefaultClient("service_foodplatform", "commodity_rpc", "dev.food_platform", instance)
	fmt.Println(config)
	var result RpcConfig
	err := GetRemoteConfig(config, &result)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestGetRemoteConfig3(t *testing.T) {
	instance := &NacosInstance{}
	config := &ClientConfig{
		Conn:                instance,
		NamespaceId:         nil,
		TimeoutMs:           nil,
		NotLoadCacheAtStart: nil,
		LogDir:              nil,
		CacheDir:            nil,
		LogLevel:            nil,
		DataId:              nil,
		Group:               nil,
	}
	var result any
	err := GetRemoteConfig(config, result)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

}

func TestConfigCLient(t *testing.T) {
	configClient, err := ConfigClient(DefaultClient("public", "mysql", "dev_food_platform", nil))
	if err != nil {
		t.Error(err)
	}
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "mysql",
		Group:  "dev_food_platform",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(config)
}

func TestRegisterCenter(t *testing.T) {
	//config := DefaultClient("service_foodplatform", "commodity_rpc", "dev.food_platform", nil)
	//service := RegisterService(config, "commodity_rpc")
	//fmt.Println(service)
	client, err := ServiceClient(DefaultClient("test", "mysql", "dev_food_platform", nil))
	if err != nil {
		t.Error(err)
	}
	instance, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8848,
		Enable:      true,
		Healthy:     true,
		Weight:      10,
		Metadata:    map[string]string{"version": "1.0"},
		ClusterName: "test",
		GroupName:   "dev.food_platform",
		Ephemeral:   true,
		ServiceName: "test_app",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(instance)
}

func TestRegisterServe(t *testing.T) {
	param := vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        8848,
		Enable:      true,
		Healthy:     true,
		Weight:      10,
		Metadata:    map[string]string{"version": "1.0"},
		ClusterName: "test",
		GroupName:   "dev_food_platform",
		Ephemeral:   true,
		ServiceName: "test_app",
	}
	service := RegisterService(DefaultClient("public", "mysql", "dev_food_platform", nil), param)
	fmt.Println(service)
}

func TestGetService(t *testing.T) {
	instance := DefaultClient("public", "mysql", "dev_food_platform", nil)
	var param = vo.GetServiceParam{
		ServiceName: "test_app",            // 服务命
		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		GroupName:   "dev_food_platform",   // 默认值DEFAULT_GROUP
	}
	service, err := GetService(instance, param)
	if err != nil {
		t.Error(err)
	}
	t.Log(service)
}

func TestGetAllServices(t *testing.T) {
	instance := DefaultClient("public", "mysql", "dev_food_platform", nil)
	param := vo.SelectAllInstancesParam{
		ServiceName: "test_app",
		GroupName:   "dev_food_platform", // 默认值DEFAULT_GROUP
		Clusters:    []string{"test"},    // 默认值DEFAULT
	}
	services, err := GetAllServices(instance, param)
	if err != nil {
		t.Error(err)
	}
	t.Log(services)
}
