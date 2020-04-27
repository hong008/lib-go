package redisUtil

func SADD(key string, members ...interface{}) error {
	conn, err := getConn()
	if err != nil {
		return err
	}
	defer conn.close()
	err = conn.sAdd(key, members...)
	return err
}

func SIsMember(key string, member interface{}) (bool, error) {
	conn, err := getConn()
	if err != nil {
		return false, err
	}
	defer conn.close()
	result, err := conn.sIsMember(key, member)
	return result, err
}

func SCard(key string) (int, error) {
	conn, err := getConn()
	if err != nil {
		return 0, err
	}
	defer conn.close()
	count, err := conn.sCARD(key)
	return count, err
}

func Smembers(key string) ([]interface{}, error) {
	conn, err := getConn()
	if err != nil {
		return nil, err
	}
	defer conn.close()
	result, err := conn.sMembers(key)
	return result, err
}
