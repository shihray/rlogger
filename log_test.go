package rlogger

import (
	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog(true, "test", "./log_file", nil)

	SetTraceSpan("trace_id", "span_id")

	Debug("test %v", "Hello")
	TWarning(nil, "trace %v", "Hi")
}
