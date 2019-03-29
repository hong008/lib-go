package redisUtils

func SADD(key string, members ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	err = conn.sAdd(key, members...)
	conn.close()
	return err
}

func SIsMember(key string, member interface{}) (bool, error) {
	conn, err := getConn()
	if err != nil {
		return false, err
	}
	result, err := conn.sIsMember(key, member)
	conn.close()
	return result, err
}

func SCard(key string) (int, error) {
	conn, err := getConn()
	if err != nil {
		return 0, err
	}
	count, err := conn.sCARD(key)
	conn.close()
	return count, err
}

func Smembers(key string) ([]interface{}, error) {
	conn, err := getConn()
	if err != nil {
		return nil, err
	}
	result, err := conn.sMembers(key)
	conn.close()


	return result, err
}
