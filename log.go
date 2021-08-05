package rlogger

import (
	beego "github.com/shihray/rlogger/beego"
)

var globalLogger *beego.BeeLogger

// InitLog 初始化日志
func InitLog(debug bool, ProcessID string, logDir string, settings map[string]interface{}) {
	globalLogger = NewBeegoLogger(debug, ProcessID, logDir, settings)
}

func GetLogger() *beego.BeeLogger {
	if globalLogger == nil {
		globalLogger = beego.NewLogger()
	}
	return globalLogger
}


func Debug(format string, a ...interface{}) {
	GetLogger().Debug(nil, format, a...)
}

func Info(format string, a ...interface{}) {
	GetLogger().Info(nil, format, a...)
}

func Error(format string, a ...interface{}) {
	GetLogger().Error(nil, format, a...)
}

func Warning(format string, a ...interface{}) {
	GetLogger().Warning(nil, format, a...)
}

func Close() {
	GetLogger().Close()
}
