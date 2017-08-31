package glog

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelFatal
	LevelNone
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

var levelName = []string{
	"DEBU",
	"INFO",
	"WARN",
	"ERRO",
	"PANI",
	"FATA",
}

func New(w io.Writer) *Logger {
	return &Logger{
		l: log.New(w, "", 0),
	}
}

func NewWithTimeFormat(w io.Writer, timeFormat string) *Logger {
	return &Logger{
		l:  log.New(w, "", 0),
		tf: timeFormat,
	}
}

type Logger struct {
	level int
	l     *log.Logger
	tf    string
}

func (self *Logger) SetFlags(flag int) *Logger {
	self.l.SetFlags(flag)
	return self
}

func (self *Logger) SetLevel(level int) *Logger {
	self.level = level
	return self
}

func (self *Logger) doLog(level int, message string) {
	if level < self.level {
		return
	}
	self.l.SetPrefix(levelName[level] + " ")
	self.l.Output(3, message)
}

func (self *Logger) Debug(message string) {
	self.doLog(LevelDebug, message)
}

func (self *Logger) Info(message string) {
	self.doLog(LevelInfo, message)
}

func (self *Logger) Warn(message string) {
	self.doLog(LevelWarn, message)
}

func (self *Logger) Error(message string) {
	self.doLog(LevelError, message)
}

func (self *Logger) Panic(message string) {
	self.doLog(LevelPanic, message)
	panic(message)
}

func (self *Logger) Fatal(message string) {
	self.doLog(LevelFatal, message)
	os.Exit(1)
}

func (self *Logger) Debugf(formatStr string, args ...interface{}) {
	self.doLog(LevelDebug, fmt.Sprintf(formatStr, args...))
}

func (self *Logger) Infof(formatStr string, args ...interface{}) {
	self.doLog(LevelInfo, fmt.Sprintf(formatStr, args...))
}

func (self *Logger) Warnf(formatStr string, args ...interface{}) {
	self.doLog(LevelWarn, fmt.Sprintf(formatStr, args...))
}

func (self *Logger) Errorf(formatStr string, args ...interface{}) {
	self.doLog(LevelError, fmt.Sprintf(formatStr, args...))
}

func (self *Logger) Panicf(formatStr string, args ...interface{}) {
	self.doLog(LevelPanic, fmt.Sprintf(formatStr, args...))
	panic(fmt.Sprintf(formatStr, args...))
}

func (self *Logger) Fatalf(formatStr string, args ...interface{}) {
	self.doLog(LevelFatal, fmt.Sprintf(formatStr, args...))
	os.Exit(1)
}
