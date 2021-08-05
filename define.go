package rlogger

import "rlogger/utils"

// TraceSpan A SpanID refers to a single span.
type TraceSpan interface {

	// Trace is the root ID of the tree that contains all of the spans
	// related to this one.
	TraceID() string

	// Span is an ID that probabilistically uniquely identifies this
	// span.
	SpanID() string

	ExtractSpan() TraceSpan
}

// TraceSpanImp TraceSpanImp
type TraceSpanImp struct {
	Trace string `json:"Trace"`
	Span  string `json:"Span"`
}


// TraceID TraceID
func (t *TraceSpanImp) TraceID() string {
	return t.Trace
}

// SpanID SpanID
func (t *TraceSpanImp) SpanID() string {
	return t.Span
}

// ExtractSpan ExtractSpan
func (t *TraceSpanImp) ExtractSpan() TraceSpan {
	return &TraceSpanImp{
		Trace: t.Trace,
		Span:  utils.GenerateID().String(),
	}
}
