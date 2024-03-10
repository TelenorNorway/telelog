package encode

import (
	"github.com/telenornorway/slf4go"
	"time"
)

type LogPayload struct {
	// About the logger

	Name  string
	Level slf4go.LogLevel

	// About the log message

	At        time.Time
	Format    string
	Arguments []any
	Fields    map[string]string

	// About the caller

	CallerFile                string
	CallerLine                int
	CallerFunctionName        string
	CallerHasFunctionReceiver bool
	CallerFunctionReceiver    string
}

type Encoder interface {
	Encode(payload *LogPayload) []byte
}
