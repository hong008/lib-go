package utils

import (
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"syscall"
	"time"

	"github.com/pyihe/util/typo"
)

var (
	mailChecker  = regexp.MustCompile(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`)
	phoneChecker = regexp.MustCompile(`^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$`)
)

type Tool interface {
	//校验邮箱格式
	CheckMailFormat(mail string) bool
	//校验电话号码格式
	CheckPhoneFormat(phone string) bool
	//内部随机生成一个数，判断是否小于per
	LessThanIn100(per int) bool
	//如果监听到系统中断信号，则执行onNotify()
	Notify(onNotify func())
	//src是否包含ele
	Contain(src interface{}, ele interface{}) bool
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

//生成一个1-100的随机数, 用于简单的判断概率
func (t *tool) LessThanIn100(per int) bool {
	if per < 1 || per > 100 {
		panic("input must between 1 and 100")
	}
	return per >= typo.RandInt(1, 100)
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
