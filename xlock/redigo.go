package xlock

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type RedisClient struct {
	redis.Conn
}

var (
	RedisConn RedisClient
)

type RedisConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewRedisConn(conf RedisConfig) {
	dial, err := redis.Dial("tcp", conf.Address)
	if err != nil {
		log.Printf("redis连接失败，err%v", err)
	}
	if conf.Password != "" {
		dial.Do("auth", conf.Password)
	}
	_, err = dial.Do("select", conf.DB)
	if err != nil {
		log.Printf("redis初始化数据库失败，err%v", err)
	}
	RedisConn = RedisClient{dial}
}

func Init(conn *redis.Conn) {
	RedisConn = RedisClient{*conn}
}
