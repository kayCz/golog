package kaylog

import (
	"log"
	"os"
	"strings"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	"fmt"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Option struct {
	Log Log
}

type Log struct {
	Level, FilePath, Flags string
}

type Logger struct {
	log *log.Logger
	level int
	flags int
}

func NewInstance(confPath string) (*Logger, error){

	l := &Logger{}
	var option Option
	e := loadToml(confPath, &option)

	if e != nil {
		return l, e
	}

	return newLogger(&option)

}

func newLogger(option *Option) (*Logger,error){
	file , err := os.OpenFile(option.Log.FilePath, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		return l, err
	}

	switch strings.ToLower(option.Log.Level) {
	case "debug": l.level = DEBUG
	case "info" : l.level = INFO
	case "warn" : l.level = WARN
	case "error": l.level = ERROR
	case "fatal": l.level = FATAL
	default:      l.level = ERROR
	}
	switch strings.ToLower(option.Log.Flags) {
	case "ldate": 		l.flags = log.Ldate
	case "llongfile" : 	l.flags = log.Llongfile
	case "lmicroseconds" :  l.flags = log.Lmicroseconds
	case "lshortfile": 	l.flags = log.Lshortfile
	case "ltime": 		l.flags = log.Ltime
	case "lutc": 		l.flags = log.LUTC
	case "lltdflags":	l.flags = log.LstdFlags
	default:      		l.flags = log.LstdFlags
	}

	logger := log.New(file, "", l.flags)
	l.log = logger

	return l, nil
}


func loadToml(path string, option *Option) error{

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	buff, er := ioutil.ReadAll(f)
	if er != nil {
		return er
	}
	e := toml.Unmarshal(buff, &option)
	if e != nil {
		return e
	}
	return nil
}


func (logger *Logger) Debug(v ...interface{})  {
	if logger.level >= DEBUG {
		logger.log.SetPrefix("[DEBUG] ")
		logger.log.Println(v)
	}
}

func (logger *Logger) Info(v ...interface{}) {
	if logger.level >= INFO {
		logger.log.SetPrefix("[INFO] ")
		logger.log.Println(v)
	}
}

func (logger *Logger) Warn(v ...interface{})  {
	if logger.level >= WARN {
		logger.log.SetPrefix("[WARN] ")
		logger.log.Println(v)
	}
}

func (logger *Logger) Error(v ...interface{})  {
	if logger.level >= ERROR {
		logger.log.SetPrefix("[ERROR] ")
		logger.log.Println(v)
	}
}

func (logger *Logger) Fatal(v ...interface{})  {
	if logger.level >= FATAL {
		logger.log.SetPrefix("[FATAL] ")
		logger.log.Println(v)
	}
}

func (logger *Logger) SetLogger(log *log.Logger) {
	logger.log = log
}

func (logger *Logger) SetLevel(level int)  {
	logger.level = level
}

func (logger *Logger) SetFlags(flags int)  {
	logger.log.SetFlags(flags)
}

func (logger *Logger) SetFilePath(path string) {
	file , err := os.OpenFile(path, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("cant open or create this file path")
		return
	}
	logger.log.SetOutput(file)

}

