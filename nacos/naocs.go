package nacos

// 默认配置
var (
	// 配置中心
	ipAddr       = "127.0.0.1"
	username     = "nacos"
	password     = "nacos"
	port         = uint64(8848)
	scheme1      = "http"
	scheme2      = "https"
	contextPath  = "/nacos"
	timeout      = uint64(5000)
	notLoadCache = true
	logDir       = "/tmp/nacos/log"
	cacheDir     = "/tmp/nacos/cache"
	logLevel     = "debug"

	// 注册中心
	weight    float64 = 10
	enable            = true
	healthy           = true
	ephemeral         = true
)

// nacos远程连接实例对象
type NacosInstance struct {
	IpAddr      *string `json:"ipaddr" yaml:"ipaddr"`
	Username    *string `json:"username" yaml:"username"`
	Password    *string `json:"password" yaml:"password"`
	Port        *uint64 `json:"port" yaml:"port"`
	Scheme      *string `json:"scheme" yaml:"scheme"`
	ContextPath *string `json:"contextPath" yaml:"contextPath"`
}

// 客户端配置对象
type ClientConfig struct {
	Conn                *NacosInstance `json:"conn" yaml:"conn"`
	NamespaceId         *string        `json:"namespaceId" yaml:"namespaceId"`
	TimeoutMs           *uint64        `json:"timeoutMs" yaml:"timeoutMs"`
	NotLoadCacheAtStart *bool          `json:"notLoadCacheAtStart" yaml:"notLoadCacheAtStart"`
	LogDir              *string        `json:"logDir" yaml:"logDir"`
	CacheDir            *string        `json:"cacheDir" yaml:"cacheDir"`
	LogLevel            *string        `json:"logLevel" yaml:"logLevel"`
	DataId              *string        `json:"dataId" yaml:"dataId"`
	Group               *string        `json:"group" yaml:"group"`
}

// 默认本地连接实例对象 http
func NewLocalInstance() *NacosInstance {
	return &NacosInstance{
		IpAddr:      &ipAddr,
		Username:    &username,
		Password:    &password,
		Port:        &port,
		Scheme:      &scheme1,
		ContextPath: &contextPath,
	}
}

// 默认远程连接实例对象 https
func NewRemoteInstance(ip string, port uint64, username, password string) *NacosInstance {
	return &NacosInstance{
		IpAddr:      &ip,
		Username:    &username,
		Password:    &password,
		Port:        &port,
		Scheme:      &scheme2,
		ContextPath: &contextPath,
	}
}

// 默认客户端配置实例
func DefaultClient(nameSpaceId string, dataId string, group string, instance *NacosInstance) *ClientConfig {
	if instance == nil {
		instance = NewLocalInstance()
	}
	return &ClientConfig{
		Conn:                instance,
		NamespaceId:         &nameSpaceId,
		TimeoutMs:           &timeout,
		NotLoadCacheAtStart: &notLoadCache,
		LogDir:              &logDir,
		CacheDir:            &cacheDir,
		LogLevel:            &logLevel,
		Group:               &group,
		DataId:              &dataId,
	}
}
