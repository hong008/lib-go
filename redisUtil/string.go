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
	keys, err = conn.getKeys(pattern)
	conn.close()
	return
}

func GetString(key string) (value string, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	value, err = conn.getString(key)
	conn.close()
	return
}

func SetString(key, value string) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.setString(key, value)
	conn.close()
	return err
}

func GetBytes(key string) (value []byte, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	value, err = conn.getBytes(key)
	conn.close()
	return value, err
}

func SetBytes(key string, value []byte) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.setBytes(key, value)
	conn.close()
	return err
}

func GetInt(key string) (value int, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	value, err = conn.getInt(key)
	conn.close()
	return value, err
}

func GetInt64(key string) (value int64, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	value, err = conn.getInt64(key)
	conn.close()
	return value, err
}

func SetInt(key string, value int64) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.setInt(key, value)
	conn.close()
	return err
}

func GetStruct(key string, data interface{}) (value interface{}, err error) {
	conn, err := getConn()
	if err != nil {
		return value, err
	}
	value, err = conn.getStruct(key, data)
	conn.close()
	return value, err
}

func SetStruct(key string, data interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.setStruct(key, data)
	conn.close()
	return err
}
