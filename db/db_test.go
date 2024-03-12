package db

import (
	"context"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {
	conf := MySQL{
		DSN: "root:123456@tcp(127.0.0.1:3306)/demo",
	}
	MysqlInit(conf)
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
		t.Error(err)
	}
	t.Log(user)
}

func TestRedis(t *testing.T) {
	conf := Redis{
		Address:  "127.0.0.1:6379",
		Username: "",
		Password: "123456",
		DB:       0,
	}
	GoRedisInit(conf)
	GoRedis.Set(context.Background(), "hello", "word", 5)
	s := GoRedis.Get(context.Background(), "hello").String()
	t.Log(s)

	//
}

func TestRediGo(t *testing.T) {
	conf := Redis{
		Address: "127.0.0.1:6379",
		DB:      0,
	}
	RedigoInit(conf)
	_, err := Redigo.Do("set", "hello", "word")
	if err != nil {
		t.Error(err)
	}
	reply, err := Redigo.Do("get", "hello")
	if err != nil {
		t.Error(err)
	}
	s := string(reply.([]byte))
	t.Log(s)
}
