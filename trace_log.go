package rlogger

import (
	beego "github.com/shihray/rlogger/beego"
	"github.com/shihray/rlogger/utils"
	"sync"
)

var (
	globalSpan TraceSpan
	setOnce    sync.Once
)

func CreateRootTrace() TraceSpan {
	return &TraceSpanImp{
		Trace: utils.GenerateID().String(),
		Span:  utils.GenerateID().String(),
	}
}

func CreateTrace(trace, span string) TraceSpan {
	return &TraceSpanImp{
		Trace: trace,
		Span:  span,
	}
}

func SetTraceSpan(trace, span string) {
	tmp := &TraceSpanImp{
		Trace: trace,
		Span:  span,
	}
	setOnce.Do(func() {
		globalSpan = tmp
	})
}

func GetTraceSpan() TraceSpan {
	return globalSpan
}

func TDebug(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		GetLogger().Debug(
			&beego.BeegoTraceSpan{
				Trace: span.TraceID(),
				Span:  span.SpanID(),
			}, format, a...)
	} else if globalSpan != nil {
		GetLogger().Debug(
			&beego.BeegoTraceSpan{
				Trace: globalSpan.TraceID(),
				Span:  globalSpan.SpanID(),
			}, format, a...)
	} else {
		GetLogger().Debug(nil, format, a...)
	}
}

func TInfo(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		GetLogger().Info(
			&beego.BeegoTraceSpan{
				Trace: span.TraceID(),
				Span:  span.SpanID(),
			}, format, a...)
	} else if globalSpan != nil {
		GetLogger().Info(
			&beego.BeegoTraceSpan{
				Trace: globalSpan.TraceID(),
				Span:  globalSpan.SpanID(),
			}, format, a...)
	} else {
		GetLogger().Info(nil, format, a...)
	}
}

func TError(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		GetLogger().Error(
			&beego.BeegoTraceSpan{
				Trace: span.TraceID(),
				Span:  span.SpanID(),
			}, format, a...)
	} else if globalSpan != nil {
		GetLogger().Error(
			&beego.BeegoTraceSpan{
				Trace: globalSpan.TraceID(),
				Span:  globalSpan.SpanID(),
			}, format, a...)
	} else {
		GetLogger().Error(nil, format, a...)
	}
}

func TWarning(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		GetLogger().Warning(
			&beego.BeegoTraceSpan{
				Trace: span.TraceID(),
				Span:  span.SpanID(),
			}, format, a...)
	} else if globalSpan != nil {
		GetLogger().Warning(
			&beego.BeegoTraceSpan{
				Trace: globalSpan.TraceID(),
				Span:  globalSpan.SpanID(),
			}, format, a...)
	} else {
		GetLogger().Warning(nil, format, a...)
	}
}
