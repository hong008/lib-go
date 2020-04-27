package redisUtil

func HGet(key string, field string) ([]byte, error) {
	conn, err := getConn()
	if err != nil {
		return []byte{}, err
	}
	defer conn.close()
	value, err := conn.hGet(key, field)
	return value, err
}

func HSet(key string, field string, value interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.hSet(key, field, value)
	return err
}

func HGetAll(key string) ([]interface{}, error) {
	conn, err := getConn()
	if err != nil {
		return nil, err
	}
	defer conn.close()
	result, err := conn.hGetAll(key)
	return result, err
}

func HKeys(key string) ([]string, error) {
	conn, err := getConn()
	if err != nil {
		return []string{}, err
	}
	defer conn.close()
	keys, err := conn.hKeys(key)
	return keys, err
}

func HMset(key string, field_value ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.hMset(key, field_value...)
	return err
}

func HDel(key, field string) (int, error) {
	conn, err := getConn()
	if err != nil {
		return 0, err
	}
	defer conn.close()
	num, err := conn.hDel(key, field)
	return num, err
}
