package db

import (
	"context"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	GoRedis *redis.Client
	Redigo  redigo.Conn
)

func MysqlInit(conf MySQL) {
	db, err := gorm.Open(mysql.Open(conf.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	DB = db
}

func GoRedisInit(conf Redis) {
	GoRedis = redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	})
	if err := GoRedis.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func RedigoInit(conf Redis) {
	dial, err := redigo.Dial("tcp", conf.Address)
	if err != nil {
		panic(err)
	}
	if conf.Password != "" {
		_, err = dial.Do("auth", conf.Password)
		if err != nil {
			panic(err)
		}
	}
	_, err = dial.Do("select", conf.DB)
	if err != nil {
		panic(err)
	}
	Redigo = dial
}
