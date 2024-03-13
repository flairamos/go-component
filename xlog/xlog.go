package xlog

import (
	"fmt"
	ct "github.com/daviddengcn/go-colortext"
	"log"
	"os"
)

var GoLogger goLogger

type goLogger struct {
	Log *log.Logger
}

func Init() {
	var logger goLogger
	logger.Log = log.Default()
	logger.Log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	logger.Log.SetPrefix("")
	GoLogger = logger
}

// 获取gopath路径
func getGoPathDir() string {
	return os.Getenv("GOPATH")
}

// SetLogPath dir路径为绝对路径，最后一个/为文件名
// dir路径为相对路径，在当前目录下，最后一个/为文件名
func SetLogPath(dir string) {
	file, err := os.OpenFile(dir+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		GoLogger.Log.Println("日志文件创建失败err：", err)
	}
	GoLogger.Log.SetOutput(file)
	return
}

// SetPrefix 设置log前缀
func SetPrefix(prefix string) {
	GoLogger.Log.SetPrefix(prefix)
}

// SetFlags 设置日志标识
func SetFlags(flag int) {
	GoLogger.Log.SetFlags(flag)
}

func init() {
	Init()
}

func (l *goLogger) InfoP(v ...any) {
	ct.Foreground(ct.White, false)
	str := fmt.Sprintln("info \t", v)
	l.Log.Output(2, str)
}

func (l *goLogger) WarnP(v ...any) {
	ct.Foreground(ct.Green, false)
	str := fmt.Sprintln("warn \t", v)
	l.Log.Output(2, str)
}

func (l *goLogger) ErrorP(v ...any) {
	ct.Foreground(ct.Yellow, false)
	str := fmt.Sprintln("error\t", v)
	l.Log.Output(2, str)
}

func (l *goLogger) FatalP(v ...any) {
	ct.Foreground(ct.Red, false)
	str := fmt.Sprintln("fatal\t", v)
	l.Log.Output(2, str)
	os.Exit(1)
}

// format和fmt.Sprintf()方法用法一致
func (l *goLogger) InfoF(format string, v ...any) {
	ct.Foreground(ct.White, false)
	str := fmt.Sprintf("info \t"+format, v...)
	l.Log.Output(2, str)
}

func (l *goLogger) WarnF(format string, v ...any) {
	ct.Foreground(ct.Green, false)
	str := fmt.Sprintf("warn \t"+format, v...)
	l.Log.Output(2, str)
}

func (l *goLogger) ErrorF(format string, v ...any) {
	ct.Foreground(ct.Yellow, false)
	str := fmt.Sprintf("error\t"+format, v...)
	l.Log.Output(2, str)
}

func (l *goLogger) FatalF(format string, v ...any) {
	ct.Foreground(ct.Red, false)
	str := fmt.Sprintf("fatal\t"+format, v...)
	l.Log.Output(2, str)
}

func Info(v ...any) {
	GoLogger.InfoP(v...)
}
func Warn(v ...any) {
	GoLogger.WarnP(v...)
}
func Error(v ...any) {
	GoLogger.ErrorP(v...)
}
func Fatal(v ...any) {
	GoLogger.FatalP(v...)
}
func InfoF(format string, v ...any) {
	GoLogger.InfoF(format, v...)
}
func WarnF(format string, v ...any) {
	GoLogger.WarnF(format, v...)
}
func ErrorF(format string, v ...any) {
	GoLogger.ErrorF(format, v...)
}
func FatalF(format string, v ...any) {
	GoLogger.FatalF(format, v...)
}
