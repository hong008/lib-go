package redisUtils

func RPush(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.rpush(key, values...)
	conn.close()
	return err
}

func RPushX(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.rpushx(key, values...)
	conn.close()
	return err
}

func LPush(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.lpush(key, values...)
	conn.close()
	return err
}

func LPushX(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.lpushx(key, values...)
	conn.close()
	return err
}

func LPop(key string) (result []byte, err error) {
	conn, err := getConn()
	if err != nil {
		return []byte{}, err
	}
	result, err = conn.lpop(key)
	conn.close()
	return result, err
}

func RPop(key string) (result []byte, err error) {
	conn, err := getConn()
	if err != nil {
		return []byte{}, err
	}
	result, err = conn.rpop(key)
	conn.close()
	return result, err
}
