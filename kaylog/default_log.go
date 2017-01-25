package kaylog

import (
	"log"
)

var defaultLog *Logger = new("./conf/log.toml")

func new(configPath string) *Logger {
	l , e := NewInstance(configPath)
	if e != nil {
		panic("no default log config")
	}
	return l
}

func Debug(v ...interface{})  {
	defaultLog.Debug(v)
}

func Info(v ...interface{}) {
	defaultLog.Info(v)
}

func Warn(v ...interface{})  {
	defaultLog.Warn(v)
}

func Error(v ...interface{})  {
	defaultLog.Error(v)
}

func Fatal(v ...interface{})  {
	defaultLog.Fatal(v)
}

func SetLogger(logger *log.Logger) {
	defaultLog.SetLogger(logger)
}

func SetLevel(level int)  {
	defaultLog.SetLevel(level)
}

func SetFlags(flags int) {
	defaultLog.SetFlags(flags)
}

func SetFilePath(path string) {
	defaultLog.SetFilePath(path)
}