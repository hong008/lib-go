package redisUtil

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
)

type myRedisConn struct {
	conn redis.Conn //redis连接池
}

func (conn *myRedisConn) Close() error {
	if conn.conn != nil {
		return conn.conn.Close()
	}
	return fmt.Errorf("no redis conn")
}

/*****************************string操作**********************************/
//获取指定模式的key
func (c *myRedisConn) getKeys(pattern string) (keys []string, err error) {
	if c.conn == nil {
		return []string{}, fmt.Errorf("unavailable conn")
	}
	if pattern == "" {
		pattern = "*"
	}

	keys, err = redis.Strings(c.conn.Do("KEYS", pattern))
	return keys, err
}

//string
func (c *myRedisConn) getString(key string) (value string, err error) {
	if c.conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.String(c.conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return value, err
	}
	return value, err
}

func (c *myRedisConn) setString(key, value string) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := redis.String(c.conn.Do("SET", key, value))
	return err
}

//[]byte
func (c *myRedisConn) getBytes(key string) (value []byte, err error) {
	if c.conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.Bytes(c.conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return value, err
}

//
func (c *myRedisConn) setBytes(key string, value []byte) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}

	_, err := redis.String(c.conn.Do("SET", key, value))
	if err != nil {
		return err
	}
	return nil
}

//int
func (c *myRedisConn) getInt(key string) (value int, err error) {
	if c.conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.Int(c.conn.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
	}
	return value, err
}

func (c *myRedisConn) getInt64(key string) (value int64, err error) {
	if c.conn == nil {
		return value, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return value, fmt.Errorf("key cannot be empty")
	}
	value, err = redis.Int64(c.conn.Do("GET", key))
	if err == redis.ErrNil {
		err = nil
	}
	return value, err
}

func (c *myRedisConn) setInt(key string, value int64) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := redis.String(c.conn.Do("SET", key, value))
	return err
}

//struct
func (c *myRedisConn) getStruct(key string, data interface{}) (result interface{}, err error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	bytes, err := redis.Bytes(c.conn.Do("GET", key))
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

func (c *myRedisConn) setStruct(key string, data interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	bytes, err := msgpack.Marshal(data)
	if err != nil {
		return err
	}
	_, err = redis.String(c.conn.Do("SET", key, bytes))
	return err
}

func (c *myRedisConn) del(key string) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := redis.String(c.conn.Do("DEL", key))
	return err
}

/******************************set操作*********************************/
//往key对应的set添加元素
func (c *myRedisConn) sAdd(key string, members ...interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := c.conn.Do("MULTI")
	if err != nil {
		return err
	}
	for _, m := range members {
		_, err = c.conn.Do("SADD", key, m)
		if err != nil {
			return err
		}
	}
	_, err = c.conn.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}

//判断元素是否为set的元素
func (c *myRedisConn) sIsMember(key string, member interface{}) (bool, error) {
	if c.conn == nil {
		return false, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return false, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Int(c.conn.Do("SISMEMBER", key, member))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
		return false, err
	}
	return result == 1, err
}

//随机从集合中获取元素
func (c *myRedisConn) sRandMember(key string, count int) (value []interface{}, err error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	if count <= 0 {
		return nil, fmt.Errorf("count must greater than 0")
	}
	value, err = redis.Values(c.conn.Do("SRANDMEMBER", key, count))
	return
}

//返回集合中的元素数量
func (c *myRedisConn) sCARD(key string) (int, error) {
	if c.conn == nil {
		return 0, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return 0, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Int(c.conn.Do("SCARD", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return result, err
}

//返回集合中的所有元素
func (c *myRedisConn) sMembers(key string) ([]interface{}, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Values(c.conn.Do("SMEMBERS", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		}
	}
	return result, err
}

/*****************************hash操作**********************************/
func (c *myRedisConn) hSet(key, field string, value interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	_, err := c.conn.Do("HSET", key, field, value)
	return err
}

//获取指定域的值
func (c *myRedisConn) hGet(key, field string) ([]byte, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	result, err := redis.Bytes(c.conn.Do("HGET", key, field))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}

//返回hash中所有的域
func (c *myRedisConn) hKeys(key string) (keys []string, err error) {
	if c.conn == nil {
		return keys, fmt.Errorf("unavailable conn")
	}
	keys, err = redis.Strings(c.conn.Do("HKEYS", key))
	return keys, err
}

//返回key对应的所有域和值
func (c *myRedisConn) hGetAll(key string) (value []interface{}, err error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	value, err = redis.Values(c.conn.Do("HGETALL", key))
	return value, err
}

//设置多对field-value
func (c *myRedisConn) hMset(key string, fields ...interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if len(fields)%2 != 0 {
		return fmt.Errorf("wrong number of arguments")
	}
	_, err := c.conn.Do("MULTI")
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
		_, err = redis.String(c.conn.Do("HMSET", key, field, value))
		if err != nil {
			break
		}
	}
	_, execErr := c.conn.Do("EXEC")
	return execErr
}

//删除
func (c *myRedisConn) hDel(key string, field string) (num int, err error) {
	if c.conn == nil {
		return 0, fmt.Errorf("unavailable conn")
	}
	num, err = redis.Int(c.conn.Do("HDEL", key, field))
	return num, err
}

/*****************************list操作**********************************/
//将value插入到list头部
func (c *myRedisConn) lpush(key string, values ...interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(c.conn.Do("LPUSH", key, v))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *myRedisConn) lpushx(key string, values ...interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(c.conn.Do("LPUSHX", key, v))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *myRedisConn) rpush(key string, values ...interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(c.conn.Do("RPUSH", key, v))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *myRedisConn) rpushx(key string, values ...interface{}) error {
	if c.conn == nil {
		return fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if len(values) > 0 {
		for _, v := range values {
			_, err := redis.String(c.conn.Do("RPUSHX", key, v))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//移除列表的头元素，及左边的那个元素
func (c *myRedisConn) lpop(key string) ([]byte, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Bytes(c.conn.Do("LPOP", key))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}

func (c *myRedisConn) rpop(key string) ([]byte, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("unavailable conn")
	}
	if key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	result, err := redis.Bytes(c.conn.Do("RPOP", key))
	if err == redis.ErrNil {
		err = nil
	}
	return result, err
}
