package logfmt

import (
	"fmt"
	logformat "github.com/go-logfmt/logfmt"
	"github.com/telenornorway/slf4go"
	"github.com/telenornorway/telelog/encode"
)

type LogFormatEncoder struct{}

type writer struct{ bytes []byte }

func (w *writer) Write(p []byte) (n int, err error) {
	w.bytes = append(w.bytes, p...)
	return len(p), nil
}

func LevelToString(level slf4go.LogLevel) string {
	switch level {
	case slf4go.Trace:
		return "TRACE"
	case slf4go.Debug:
		return "DEBUG"
	case slf4go.Info:
		return "INFO"
	case slf4go.Warn:
		return "WARN"
	case slf4go.Error:
		return "ERROR"
	case slf4go.Fatal:
		return "FATAL"
	}
	panic("unknown log level")
}

func (l LogFormatEncoder) Encode(payload *encode.LogPayload) []byte {
	var writer = &writer{}
	var encoder = logformat.NewEncoder(writer)
	encoder.EncodeKeyval("time", payload.At.String())
	encoder.EncodeKeyval("logger", payload.Name)
	encoder.EncodeKeyval("level", LevelToString(payload.Level))
	encoder.EncodeKeyval("message", fmt.Sprintf(payload.Format, payload.Arguments...))
	for k, v := range payload.Fields {
		encoder.EncodeKeyval(k, v)
	}
	encoder.EndRecord()
	return writer.bytes
}

func New() encode.Encoder {
	return &LogFormatEncoder{}
}
