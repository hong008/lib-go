package commonUtil

import (
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"
	"unsafe"
)

//校验邮箱格式
func CheckMailFormat(mail string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`).MatchString(mail)
}

//校验电话号码格式
func CheckPhoneFormat(phone string) bool {
	return regexp.MustCompile(`^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$`).MatchString(phone)
}

//生成min-max之间的一个随机数
func RandBetween(min, max int) int {
	if min >= max {
		panic("min bigger than max")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

//生成一个1-100的随机数, 用于简单的判断概率
func LessThanIn100(per int) bool {
	if per < 1 || per > 100 {
		panic("input must between 1 and 100")
	}
	return per >= RandBetween(1, 100)
}

//[]byte convert to string
func Bytes2String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

//string convert to []byte
func String2Bytes(s string) []byte {
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

func Notify(onNotify func()) {
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
