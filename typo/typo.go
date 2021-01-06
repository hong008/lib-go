package typo

import (
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var (
	src = rand.NewSource(time.Now().UnixNano())
)

func Bool(b bool) *bool {
	var ptr = new(bool)
	*ptr = b
	return ptr
}

func String(s string) *string {
	var ptr = new(string)
	*ptr = s
	return ptr
}

func Byte(b byte) *byte {
	var ptr = new(byte)
	*ptr = b
	return ptr
}

func Int(i int) *int {
	var ptr = new(int)
	*ptr = i
	return ptr
}

func Int32(i int32) *int32 {
	var ptr = new(int32)
	*ptr = i
	return ptr
}

func Int64(i int64) *int64 {
	var ptr = new(int64)
	*ptr = i
	return ptr
}

func Float32(f float32) *float32 {
	var ptr = new(float32)
	*ptr = f
	return ptr
}

func Float64(f float64) *float64 {
	var ptr = new(float64)
	*ptr = f
	return ptr
}

func String2Bytes(s string) []byte {
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//随机指定长度的字符串
func RandString(n int) string {
	if n <= 0 {
		panic("invalid n.")
	}
	sb := strings.Builder{}
	sb.Grow(n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

//生成min-max之间的一个随机数
func RandInt(min, max int) int {
	if min >= max {
		panic("min bigger than max")
	}
	return rand.Intn(max-min+1) + min
}
