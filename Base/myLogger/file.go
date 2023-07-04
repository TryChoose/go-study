package myLogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFile     *os.File
	maxFileSize int64
}

// NewFileLogger构造函数
func NewFileLogger(level string, filePath string, fileName string, maxFileSize int64) *FileLogger {
	logLevel, err := parseLogLevel(level)

	if err != nil {
		panic(err)
	}
	fl := &FileLogger{Level: logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) CheckSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file Info failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed ,err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s.back%s", logName, nowStr)
	//1.关闭当前的日志文件
	file.Close()
	//2.备份一下 rename
	os.Rename(logName, newLogName)
	//3.打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed err:%v", err)
		return nil, err
	}
	//4.将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}
func (f *FileLogger) initFile() error {
	filepath := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0555)
	if err != nil {
		fmt.Printf("OPEN LOG FILE FAILED,err :%v", err)
		return err
	}
	//defer fileObj.Close()
	errFileObj, err := os.OpenFile(filepath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0555)
	if err != nil {
		fmt.Printf("OPEN ERR LOG FILE FAILED,err :%v", err)
		return err
	}
	//defer errFileObj.Close()
	f.fileObj = fileObj
	f.errFile = errFileObj
	return nil
}

// 开关
func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineno := getInfo(3)
		if f.CheckSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineno, msg)
		if lv >= ERROR {
			if f.CheckSize(f.errFile) {
				newFile, err := f.splitFile(f.errFile)
				if err != nil {
					return
				}
				f.errFile = newFile
			}
			//如果记录的日志大于等于ERROR级别，还要在err日志文件中记录一下
			fmt.Fprintf(f.errFile, "[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineno, msg)
		}
	}
}

// Debug
func (f FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Trace
func (f FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Info
func (f FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning
func (f FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error
func (f FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal
func (f FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFile.Close()
}
