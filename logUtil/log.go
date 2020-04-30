package logUtil

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 34

	dTypeStd  = 0
	dTypeFile = 1

	levelT = "[T] "
	levelE = "[E] "
	levelW = "[W] "

	fileSize       = 60 * 1024 * 1024
	logPath        = "log/"
	defaultLogName = "default.log"
)

var (
	defaultLogger = &myLog{}
)

type myLog struct {
	sync.Once
	logChan    chan *data //日志写通道
	fileWriter *os.File   //输出到文件
	stdWriter  io.Writer  //输出到控制台
}

type data struct {
	dType   uint8
	content string
}

func InitLogger() {
	defaultLogger.Do(func() {
		defaultLogger = new(myLog)
		defaultLogger.logChan = make(chan *data, 1024)
		defaultLogger.init()
	})
}

func (m *myLog) init() {
	if !isExist(logPath) {
		if err := os.Mkdir(logPath, 0777); err != nil {
			panic(err)
		}
	}
	fileName := logPath + defaultLogName
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	m.fileWriter = file
	m.stdWriter = os.Stdout
	go m.writeLog()
}

func (m *myLog) writeLog() {
	if m.logChan == nil {
		return
	}
	for data := range m.logChan {
		if data == nil {
			continue
		}
		m.checkLogSize()
		if data.dType == dTypeStd {
			fmt.Fprintf(m.stdWriter, data.content+"\n")
		}
		if data.dType == dTypeFile {
			fmt.Fprintf(m.fileWriter, data.content+"\n")
		}
	}
}

func (m *myLog) checkLogSize() {
	if m.fileWriter == nil {
		return
	}
	fileInfo, err := m.fileWriter.Stat()
	if err != nil {
		panic(err)
	}
	if fileSize > fileInfo.Size() {
		return
	}
	//需要分割
	name := logPath + time.Now().Format("2006_01_02_15:04:03") + ".log"
	m.fileWriter.Close()

	err = os.Rename(logPath+defaultLogName, name)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(logPath+defaultLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	m.fileWriter = file
	return
}

////Info
func T(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	codeLine := "[" + shortFileName(file) + ":" + strconv.Itoa(line) + "]"
	content := levelT + codeLine + fmt.Sprintf(format, v...)
	var d = &data{
		dType:   dTypeStd,
		content: setColor(content, colorBlue),
	}
	defaultLogger.logChan <- d
	var d1 = &data{
		dType:   dTypeFile,
		content: content,
	}
	defaultLogger.logChan <- d1
}

//
////Error
func E(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	codeLine := "[" + shortFileName(file) + ":" + strconv.Itoa(line) + "]"
	content := levelE + codeLine + fmt.Sprintf(format, v...)

	var d = &data{
		dType:   dTypeStd,
		content: setColor(content, colorRed),
	}
	defaultLogger.logChan <- d
	var d1 = &data{
		dType:   dTypeFile,
		content: content,
	}
	defaultLogger.logChan <- d1
}

//Warn
func W(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	codeLine := "[" + shortFileName(file) + ":" + strconv.Itoa(line) + "]"
	content := levelW + codeLine + fmt.Sprintf(format, v...)

	var d = &data{
		dType:   dTypeStd,
		content: setColor(content, colorYellow),
	}
	defaultLogger.logChan <- d
	var d1 = &data{
		dType:   dTypeFile,
		content: content,
	}
	defaultLogger.logChan <- d1
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func shortFileName(file string) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}

func setColor(msg string, text int) string {
	return fmt.Sprintf("%c[%dm%s%c[0m", 0x1B, text, msg, 0x1B)
}
