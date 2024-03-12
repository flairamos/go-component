package xlock

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
)

func Lock(name string) (error, *string) {
	apply, err := RedisConn.Do("GET", name)
	if apply != nil || err != nil {
		return errors.New("locking"), nil
	}
	uniqueId, err := uuid.NewV4()
	if err != nil {
		return errors.New("uuid err"), nil
	}
	uid := uniqueId.String()
	fmt.Printf("name: %v, uid: %v \n", name, uid)
	_, err = RedisConn.Do("SET", name, uid, "EX", 10)
	if err != nil {
		RedisConn.Do("DEL", name)
		return errors.New("redis set err"), nil
	}
	return nil, &uid
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
