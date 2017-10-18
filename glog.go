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

type Logger struct {
	level int
	l     *log.Logger
}

func (self *Logger) SetFlags(flag int) *Logger {
	self.l.SetFlags(flag)
	return self
}

func (self *Logger) SetLevel(level int) *Logger {
	self.level = level
	return self
}

func (self *Logger) doLog(level int, v ...interface{}) {
	if level < self.level {
		return
	}
	self.l.SetPrefix(levelName[level] + " ")
	self.l.Output(3, fmt.Sprintln(v...))
}

func (self *Logger) Debug(v ...interface{}) {
	self.doLog(LevelDebug, v...)
}

func (self *Logger) Info(v ...interface{}) {
	self.doLog(LevelInfo, v...)
}

func (self *Logger) Warn(v ...interface{}) {
	self.doLog(LevelWarn, v...)
}

func (self *Logger) Error(v ...interface{}) {
	self.doLog(LevelError, v...)
}

func (self *Logger) Panic(v ...interface{}) {
	self.doLog(LevelPanic, v...)
	panic(fmt.Sprintln(v...))
}

func (self *Logger) Fatal(v ...interface{}) {
	self.doLog(LevelFatal, v...)
	os.Exit(1)
}

func (self *Logger) Debugf(format string, v ...interface{}) {
	self.doLog(LevelDebug, fmt.Sprintf(format, v...))
}

func (self *Logger) Infof(format string, v ...interface{}) {
	self.doLog(LevelInfo, fmt.Sprintf(format, v...))
}

func (self *Logger) Warnf(format string, v ...interface{}) {
	self.doLog(LevelWarn, fmt.Sprintf(format, v...))
}

func (self *Logger) Errorf(format string, v ...interface{}) {
	self.doLog(LevelError, fmt.Sprintf(format, v...))
}

func (self *Logger) Panicf(format string, v ...interface{}) {
	self.doLog(LevelPanic, fmt.Sprintf(format, v...))
	panic(fmt.Sprintf(format, v...))
}

func (self *Logger) Fatalf(format string, v ...interface{}) {
	self.doLog(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}
