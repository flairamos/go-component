package xlock

import (
	"fmt"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	NewRedisConn(RedisConfig{
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
	err, s := Lock("test", time.Second*10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("A is running with uuid :", s)
	err = Unlock("test", s)
	if err != nil {
		fmt.Println(err)
		return
	}
}
