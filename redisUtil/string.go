package redisUtil

import (
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	defaultPool *myPool
)

type myPool struct {
	sync.Once
	p *redis.Pool
}

//参数顺序：network/addr/password
func InitPool(redisNetwork, redisAddr string, options ...redis.DialOption) {
	defaultPool = &myPool{}
	defaultPool.Do(func() {
		defaultPool.p = &redis.Pool{
			Dial: func() (conn redis.Conn, e error) {
				return redis.Dial(redisNetwork, redisAddr, options...)
			},
			MaxIdle:     10,
			MaxActive:   0,
			IdleTimeout: 120 * time.Second,
			Wait:        true,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	})
}

func GetKeys(pattern string) (keys []string, err error) {
	conn, err := getConn()
	if err != nil {
		return keys, err
	}
	defer conn.close()
	keys, err = conn.getKeys(pattern)
	return
}

func GetString(key string) (value string, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	defer conn.close()
	value, err = conn.getString(key)
	return
}

func SetString(key, value string) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.setString(key, value)
	return err
}

func GetBytes(key string) (value []byte, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	defer conn.close()
	value, err = conn.getBytes(key)
	return value, err
}

func SetBytes(key string, value []byte) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.setBytes(key, value)
	return err
}

func GetInt(key string) (value int, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	defer conn.close()
	value, err = conn.getInt(key)
	return value, err
}

func GetInt64(key string) (value int64, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	defer conn.close()
	value, err = conn.getInt64(key)
	return value, err
}

func SetInt(key string, value int64) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.setInt(key, value)
	return err
}

func GetStruct(key string, data interface{}) (value interface{}, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	defer conn.close()
	value, err = conn.getStruct(key, data)
	return value, err
}

func SetStruct(key string, data interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.setStruct(key, data)
	return err
}
