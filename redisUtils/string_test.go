package redisUtils

import (
	"fmt"
	"testing"
)

/*
    @Create by GoLand
    @Author: hong
    @Time: 2019-03-27 17:58 
    @File: utils_test.go    
*/

func init() {
	InitRedisPool("tcp", "127.0.0.1:6379", "", 0)
}

func TestInitRedisPool(t *testing.T) {
	keys, err := GetKeys("test")
	fmt.Printf("keys = %v  err = [%v]", keys, err)
}
