package rlogger

import (
	beego "rlogger/beego"
	"rlogger/utils"
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

func TDebug(span TraceSpan, format string, a ...interface{}) {
	if span != nil {
		GetLogger().Debug(
			&beego.BeegoTraceSpan{
				Trace: span.TraceID(),
				Span:  span.SpanID(),
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
	} else {
		GetLogger().Warning(nil, format, a...)
	}
}
