package myLogger

import (
	"fmt"
	"time"
)

// 日志的结构体
type ConsoleLogger struct {
	Level LogLevel
}

// 构造方法
func NewLogger(levelStr string) *ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{Level: level}
}

// 开关
func (c ConsoleLogger) enable(level LogLevel) bool {
	return level >= c.Level
}
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineno := getInfo(3)
		fmt.Printf("[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineno, msg)
	}
}

// Debug
func (c ConsoleLogger) Debug(format string, a ...interface{}) {

	c.log(DEBUG, format, a...)

}

// Trace
func (c ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log(TRACE, format, a...)

}

// Info
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)

}

// Warning
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)

}

// Error
func (c ConsoleLogger) Error(format string, a ...interface{}) {

	c.log(ERROR, format, a...)
}

// Fatal
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)

}
