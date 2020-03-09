package redisUtils

func HGet(key string, field string) ([]byte, error) {
	conn, err := getConn()
	if err != nil {
		return []byte{}, err
	}
	value, err := conn.hGet(key, field)
	conn.close()
	return value, err
}

func HSet(key string, field string, value interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.hSet(key, field, value)
	conn.close()
	return err
}

func HGetAll(key string) ([]interface{}, error) {
	conn, err := getConn()
	if err != nil {
		return nil, err
	}
	result, err := conn.hGetAll(key)
	conn.close()
	return result, err
}

func HKeys(key string) ([]string, error) {
	conn, err := getConn()
	if err != nil {
		return []string{}, err
	}
	keys, err := conn.hKeys(key)
	conn.close()
	return keys, err
}

func HMset(key string, field_value ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.hMset(key, field_value...)
	conn.close()
	return err
}

func HDel(key, field string) (int, error) {
	conn, err := getConn()
	if err != nil {
		return 0, err
	}
	num, err := conn.hDel(key, field)
	conn.close()
	return num, err
}
