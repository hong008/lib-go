package redisUtil

func RPush(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.rpush(key, values...)
	return err
}

func RPushX(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.rpushx(key, values...)
	return err
}

func LPush(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.lpush(key, values...)
	return err
}

func LPushX(key string, values ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.lpushx(key, values...)
	return err
}

func LPop(key string) (result []byte, err error) {
	conn, err := getConn()
	if err != nil {
		return []byte{}, err
	}
	defer conn.close()
	result, err = conn.lpop(key)
	return result, err
}

func RPop(key string) (result []byte, err error) {
	conn, err := getConn()
	if err != nil {
		return []byte{}, err
	}
	defer conn.close()
	result, err = conn.rpop(key)
	return result, err
}
