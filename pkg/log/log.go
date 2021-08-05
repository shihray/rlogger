package log

import (
	beegolog "rlogger/pkg/log/beego"
	"rlogger/utils"
)

var beego *beegolog.BeeLogger
var bi *beegolog.BeeLogger

// InitLog 初始化日志
func InitLog(debug bool, ProcessID string, Logdir string, settings map[string]interface{}) {
	beego = NewBeegoLogger(debug, ProcessID, Logdir, settings)
}

// InitBI 初始化BI日志
func InitBI(debug bool, ProcessID string, Logdir string, settings map[string]interface{}) {
	bi = NewBeegoLogger(debug, ProcessID, Logdir, settings)
}

// LogBeego LogBeego
func LogBeego() *beegolog.BeeLogger {
	if beego == nil {
		beego = beegolog.NewLogger()
	}
	return beego
}

// BiBeego BiBeego
func BiBeego() *beegolog.BeeLogger {
	return bi
}

// CreateRootTrace CreateRootTrace
func CreateRootTrace() TraceSpan {
	return &TraceSpanImp{
		Trace: utils.GenerateID().String(),
		Span:  utils.GenerateID().String(),
	}
}

// CreateTrace CreateTrace
func CreateTrace(trace, span string) TraceSpan {
	return &TraceSpanImp{
		Trace: trace,
		Span:  span,
	}
}

// BiReport BiReport
func BiReport(msg string) {
	//gLogger.doPrintf(debugLevel, printDebugLevel, format, a...)
	l := BiBeego()
	if l != nil {
		l.BiReport(msg)
	}
}

// Debug Debug
func Debug(format string, a ...interface{}) {
	//gLogger.doPrintf(debugLevel, printDebugLevel, format, a...)
	LogBeego().Debug(nil, format, a...)
}

// Info Info
func Info(format string, a ...interface{}) {
	//gLogger.doPrintf(releaseLevel, printReleaseLevel, format, a...)
	LogBeego().Info(nil, format, a...)
}

// Error Error
func Error(format string, a ...interface{}) {
	//gLogger.doPrintf(errorLevel, printErrorLevel, format, a...)
	LogBeego().Error(nil, format, a...)
}

// Warning Warning
func Warning(format string, a ...interface{}) {
	//gLogger.doPrintf(fatalLevel, printFatalLevel, format, a...)
	LogBeego().Warning(nil, format, a...)
}

// TDebug TDebug
func TDebug(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Debug(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Debug(nil, format, a...)
	}
}

// TInfo TInfo
func TInfo(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Info(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Info(nil, format, a...)
	}
}

// TError TError
func TError(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Error(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Error(nil, format, a...)
	}
}

// TWarning TWarning
func TWarning(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		LogBeego().Warning(
			&beegolog.BeegoTraceSpan{
				Trace: span.TraceId(),
				Span:  span.SpanId(),
			}, format, a...)
	} else {
		LogBeego().Warning(nil, format, a...)
	}
}

// Close Close
func Close() {
	LogBeego().Close()
}
