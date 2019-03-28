package redisUtils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
)

var (
	pool *redis.Pool
)

func getConn() (*RedisClient, error) {
	c := &RedisClient{}
	if err := c.init(); err != nil {
		return c, fmt.Errorf("cannot establish conn")
	}
	return c, nil
}

type RedisClient struct {
	redis.Conn //redis连接池
}

func (c *RedisClient) init() error {
	conn := pool.Get()
	if conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	_, err := conn.Do("PING")
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

func (c *RedisClient) close() error {
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return fmt.Errorf("no redis conn")
}

/*****************************string操作**********************************/
//获取指定模式的key
func (c *RedisClient) getKeys(pattern string) (keys []string, err error) {
	if c.Conn == nil {
		return []string{}, fmt.Errorf("unavailable conn")
	}
	if pattern == "" {
		pattern = "*"
	}

	keys, err = redis.Strings(c.Do("KEYS", pattern))
	return keys, err
}

//string
func (c *RedisClient) getString(key string) (value string, err error) {
	if c.Conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.String(c.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return value, err
	}
	return value, err
}

func (c *RedisClient) setString(key, value string) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := redis.String(c.Do("SET", key, value))
	return err
}

//[]byte
func (c *RedisClient) getBytes(key string) (value []byte, err error) {
	if c.Conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.Bytes(c.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return value, err
}

//
func (c *RedisClient) setBytes(key string, value []byte) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	_, err := redis.String(c.Do("SET", key, value))
	if err != nil {
		return err
	}
	return nil
}

//int
func (c *RedisClient) getInt(key string) (value int, err error) {
	if c.Conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.Int(c.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
	}
	return value, err
}

func (c *RedisClient) getInt64(key string) (value int64, err error) {
	if c.Conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.Int64(c.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
	}
	return value, err
}

func (c *RedisClient) setInt(key string, value int64) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := redis.String(c.Do("SET", key, value))
	return err
}

//struct
func (c *RedisClient) getStruct(key string, data interface{}) (result interface{}, err error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	bytes, err := redis.Bytes(c.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	err = msgpack.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *RedisClient) setStruct(key string, data interface{}) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	bytes, err := msgpack.Marshal(data)
	if err != nil {
		return err
	}
	_, err = redis.String(c.Do("SET", key, bytes))
	return err
}

func (c *RedisClient) del(key string) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := redis.String(c.Do("DEL", key))
	return err
}

/******************************set操作*********************************/
//往key对应的set添加元素
func (c *RedisClient) sAdd(key string, members ...interface{}) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := c.Do("MULTI")
	if err != nil {
		return err
	}
	for _, m := range members {
		_, err = c.Do("SADD", key, m)
		if err != nil {
			return err
		}
	}
	_, err = c.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}

//判断元素是否为set的元素
func (c *RedisClient) sIsMember(key string, member interface{}) (bool, error) {
	if c.Conn == nil {
		return false, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return false, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Int(c.Do("SISMEMBER", key, member))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return false, err
	}
	return result == 1, err
}

//随机从集合中获取元素
func (c *RedisClient) sRandMember(key string, count int) (value []interface{}, err error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	if count <= 0 {
		return nil, fmt.Errorf("count must greater than 0")
	}
	value, err = redis.Values(c.Do("SRANDMEMBER", key, count))
	return
}

//返回集合中的元素数量
func (c *RedisClient) sCARD(key string) (int, error) {
	if c.Conn == nil {
		return 0, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return 0, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Int(c.Do("SCARD", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return result, err
}

//返回集合中的所有元素
func (c *RedisClient) sMembers(key string) ([]interface{}, error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Values(c.Do("SMEMBERS", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return result, err
}

/*****************************hash操作**********************************/
func (c *RedisClient) hSet(key, field string, value interface{}) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := c.Do("HSET", key, field, value)
	return err
}

//获取指定域的值
func (c *RedisClient) hGet(key, field string) ([]byte, error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	result, err := redis.Bytes(c.Do("HGET", key, field))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}

//返回hash中所有的域
func (c *RedisClient) hKeys(key string) (keys []string, err error) {
	if c.Conn == nil {
		return keys, fmt.Errorf("unavailable conn")
	}
	keys, err = redis.Strings(c.Do("HKEYS", key))
	return keys, err
}

//返回key对应的所有域和值
func (c *RedisClient) hGetAll(key string) (value []interface{}, err error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	value, err = redis.Values(c.Do("HGETALL", key))
	return value, err
}

//设置多对field-value
func (c *RedisClient) hMset(key string, fields ...interface{}) error {
	if c.Conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if len(fields)%2 != 0 {
		return fmt.Errorf("wrong number of arguments")
	}
	_, err := c.Do("MULTI")
	if err != nil {
		return err
	}
	var field, value interface{}
	for i, v := range fields {
		if i%2 != 0 {
			value = v
		} else {
			field = v
		}
		_, err = redis.String(c.Do("HMSET", key, field, value))
		if err != nil {
			break
		}
	}
	_, execErr := c.Do("EXEC")
	return execErr
}

//删除
func (c *RedisClient) hDel(key string, field string) (num int, err error) {
	if c.Conn == nil {
		return 0, fmt.Errorf("unavailable conn")
	}
	num, err = redis.Int(c.Do("HDEL", key, field))
	return num, err
}

/*****************************list操作**********************************/
//TODO
