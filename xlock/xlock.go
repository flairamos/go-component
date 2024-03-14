package xlock

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"time"
)

func Lock(name string, duration time.Duration) (error, string) {
	apply, err := RedisConn.Do("GET", name)
	if apply != nil || err != nil {
		return errors.New("locking"), ""
	}
	uniqueId, err := uuid.NewV4()
	if err != nil {
		return errors.New("uuid err"), ""
	}
	uid := uniqueId.String()
	_, err = RedisConn.Do("SET", name, uid, "EX", duration)
	if err != nil {
		RedisConn.Do("DEL", name)
		return errors.New("redis set err"), ""
	}
	return nil, uid
}

func Unlock(name string, uid string) error {
	reply, err := RedisConn.Do("GET", name)
	if reply == nil && err == nil {
		return nil
	}
	s := string(reply.([]byte))
	if s != uid {
		RedisConn.Do("DEL", name)
		return errors.New("can not unlock")
	}
	_, err = RedisConn.Do("DEL", name)
	if err != nil {
		RedisConn.Do("DEL", name)
		return err
	}
	return nil
}

func HardLock(name string) (error, string) {
	apply, err := RedisConn.Do("GET", name)
	if apply != nil || err != nil {
		return errors.New("locking"), ""
	}
	uniqueId, err := uuid.NewV4()
	if err != nil {
		return errors.New("uuid err"), ""
	}
	uid := uniqueId.String()
	fmt.Printf("name: %v, uid: %v \n", name, uid)
	_, err = RedisConn.Do("SET", name, uid)
	if err != nil {
		RedisConn.Do("DEL", name)
		return errors.New("redis set err"), ""
	}
	return nil, uid
}

func FlashLock(name string, duration time.Duration) bool {
	_, err := RedisConn.Do("EXPIRE", name, duration)
	if err != nil {
		RedisConn.Do("DEL", name)
		return false
	}
	return true
}
