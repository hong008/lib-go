package common

import (
	"encoding/base64"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var (
	src          = rand.NewSource(time.Now().UnixNano())
	mailChecker  = regexp.MustCompile(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`)
	phoneChecker = regexp.MustCompile(`^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$`)
)

type Tool interface {
	//校验邮箱格式
	CheckMailFormat(mail string) bool
	//校验电话号码格式
	CheckPhoneFormat(phone string) bool
	//在min和max之间生成一个随机数
	RandBetween(min, max int) int
	//内部随机生成一个数，判断是否小于per
	LessThanIn100(per int) bool
	//[]byte转string
	Bytes2String(bytes []byte) string
	//string转[]byte
	String2Bytes(s string) []byte
	//如果监听到系统中断信号，则执行onNotify()
	Notify(onNotify func())
	//随机字符串
	RandString(n int) string
	//src是否包含ele
	Contain(src interface{}, ele interface{}) bool
	//bcrypt算法加密明文密钥
	BcEncryptPass(plainPass string) (string, error)
	//bcrypt算法比较明文和密文密钥
	BcComparePass(hashPass, plainPass string) error
	//scrypt算法加密明文密钥
	ScryptPass(plainPass, salt string) (string, error)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type tool struct {
}

func NewTool() Tool {
	return new(tool)
}

//校验邮箱格式
func (t *tool) CheckMailFormat(mail string) bool {
	return mailChecker.MatchString(mail)
}

//校验电话号码格式
func (t *tool) CheckPhoneFormat(phone string) bool {
	return phoneChecker.MatchString(phone)
}

//生成min-max之间的一个随机数
func (t *tool) RandBetween(min, max int) int {
	if min >= max {
		panic("min bigger than max")
	}
	return rand.Intn(max-min+1) + min
}

//生成一个1-100的随机数, 用于简单的判断概率
func (t *tool) LessThanIn100(per int) bool {
	if per < 1 || per > 100 {
		panic("input must between 1 and 100")
	}
	return per >= t.RandBetween(1, 100)
}

//[]byte convert to string
func (t *tool) Bytes2String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

//string convert to []byte
func (t *tool) String2Bytes(s string) []byte {
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

func (t *tool) Notify(onNotify func()) {
	//SIGHUP		终端控制进程结束(终端连接断开)
	//SIGQUIT		用户发送QUIT字符(Ctrl+/)触发
	//SIGTERM		结束程序(可以被捕获、阻塞或忽略)
	//SIGINT		用户发送INTR字符(Ctrl+C)触发
	//SIGSTOP		停止进程(不能被捕获、阻塞或忽略)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGHUP:
			if onNotify == nil {
				return
			}
			onNotify()
		default:
			return
		}
	}
}

//随机指定长度的字符串
func (t *tool) RandString(n int) string {
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

//判断src中是否有元素ele
func (t *tool) Contain(src interface{}, ele interface{}) bool {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(src)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(ele, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

//用于加密web的密码
func (t *tool) BcEncryptPass(plainPass string) (string, error) {
	var encryptPass string
	data, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		return encryptPass, err
	}
	encryptPass = base64.StdEncoding.EncodeToString(data)
	return encryptPass, nil
}

//比较密码是否匹配
func (t *tool) BcComparePass(hashPass, plainPass string) error {
	hashBytes, err := base64.StdEncoding.DecodeString(hashPass)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hashBytes, []byte(plainPass))
}

func (t *tool) ScryptPass(plainPass, salt string) (string, error) {
	data, err := scrypt.Key([]byte(plainPass), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	encryptPass := base64.StdEncoding.EncodeToString(data)
	return encryptPass, nil
}
