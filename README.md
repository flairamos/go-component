# go-component
go语言实用小组件

## 安装

```shell
go get github.com/flairamos/go-component/xxx
```

## 用法

#### db

动态创建数据库连接实例

```go
import "github.com/flairamos/go-component/db"
func main(){
	/**
	mysql
	 */
    conf := db.MySQL{
        DSN: "root:123456@tcp(127.0.0.1:3306)/demo",
    }
    db.MysqlInit(conf)
    type User struct {
        Id          int       `gorm:"id"`
        Username    string    `gorm:"username"`
        Password    string    `gorm:"password"`
        PhoneNumber string    `gorm:"phone_number"`
        Role        string    `gorm:"role"`
        Time        time.Time `gorm:"xtime"`
    }
    var user User
    err := DB.First(&user).Error
    if err != nil {
        panic(err)
    }
    fmt.Println(user)
	
	/**
	redis
	 */
    conf := Redis{
        Address: "127.0.0.1:6379",
        DB:      0,
    }
    RedigoInit(conf)
    _, err := Redigo.Do("set", "hello", "word")
    if err != nil {
        panic(err)
    }
    reply, err := Redigo.Do("get", "hello")
    if err != nil {
        panic(err)
    }
    s := string(reply.([]byte))
}
```

> mysql使用的gorm框架，结构体需要使用gorm的tag
> redis使用redigo框架,返回值是字节数组

#### nacos

nacos配置中心

```go
// 获取配置
import "github.com/flairamos/go-component/nacos"

func main(){
    instance := NewRemoteInstance("nacos.nacostest.xyz", 443, "nacos", "123456")
    fmt.Println(instance)
    config := DefaultClient("service_foodplatform", "commodity_rpc", "dev.food_platform", instance)
    fmt.Println(config)
    var result RpcConfig
    err := GetRemoteConfig(config, &result)
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}

// 跟多操作使用ConfigClient[config_client.IConfigClient]方法
// https://github.com/nacos-group/nacos-sdk-go
```
nacos注册中心

```go
import (
	"github.com/flairamos/go-component/nacos"
    "github.com/nacos-group/nacos-sdk-go/vo"
)
func main(){
	/**
	注册服务
	 */
    param := vo.RegisterInstanceParam{
        Ip:          "127.0.0.1",   // 服务实例IP
        Port:        8848,          // 服务实例端口
        Enable:      true,          // 是否上线
        Healthy:     true,          // 是否健康
        Weight:      10,            // 权重
        Metadata:    map[string]string{"version": "1.0"},   // 扩展信息
        ClusterName: "test",        // 集群名
        GroupName:   "dev.food_platform",                   // 分组名
        Ephemeral:   true,          // 是否临时实例
        ServiceName: "test_app",    // 服务名
    }
    service := RegisterService(DefaultClient("public", "mysql", "dev_food_platform", nil), param)
	fmt.Println(service)
	
	
	/**
	获取服务列表
	 */
    instance := DefaultClient("public", "mysql", "dev_food_platform", nil)
    param := vo.SelectAllInstancesParam{
        ServiceName: "test_app",          // 服务名
        GroupName:   "dev_food_platform", // GroupID（默认值DEFAULT_GROUP）
        Clusters:    []string{"test"},    // 集群名列表（默认值DEFAULT）
    }
    services, err := GetAllServices(instance, param)
    if err != nil {
        t.Error(err)
    }
    t.Log(services)
}

// 更多自定义操作使用ServiceClient[naming_client.INamingClient]方法
// https://github.com/nacos-group/nacos-sdk-go
```

### xjson

json解析

```go
// json字符串转化为go切片
Json2List[T interface{}](jsonData string) (*[]T, error)
```

```go
// 解析json字符串为结构体
Json2Struct[T interface{}](jsonData string) (*T, error)
```

```go
// 解析json字符串为结构体切片
Json2StructList[T interface{}](jsonData string) (*[]T, error)
```


#### pin

返回数据指针

|方法|说明|
|:----:|:----:|
|I64(i64 int64)|返回int64指针|
|tr(str string)|返回string指针|
|Any[T any](t T) *T|返回任意类型指针|

#### xlog

日志库go自带日志包装

|        日志级别        |           说明           |
|:------------------:|:----------------------:|
|        Info        | 基本信息P后缀为标准输出，F后缀为格式化输出 |
|        Warn        | 警告信息P后缀为标准输出，F后缀为格式化输出 |           
| Error |        错误打印，不会退出程序，P后缀为标准输出，F后缀为格式化输出 |       
|        Fatal        |        错误打印，退出程序，P后缀为标准输出，F后缀为格式化输出 |

```go
import "github.com/flairamos/go-component/xlog"

func main() {
    xlog.GoLogger.InfoP("hello")
	xlog.GoLogger.InfoF("hello %s", "world")
}
```



#### xtime

时间解析库


#### xlock

redis锁

在方法内部上锁使方法执行一次，方法执行结束后解锁。默认10秒超时

```go
import "github.com/flairamos/go-component/xlock"
func main() {
    xNewRedisConn(RedisConfig{
        Address:  "127.0.0.1:6379",
        Username: "",
        Password: "",
        DB:       0,
    })
    A()
    go A()
    go A()
    time.Sleep(time.Second * 5)
	
}
func A() {
    err, s := Lock("test")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("A is running")
    err = Unlock("test", *s)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```





