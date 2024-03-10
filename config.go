package telelog

import (
	"github.com/telenornorway/slf4go"
	"github.com/telenornorway/telelog/transport"
)

type Config struct {
	Level      slf4go.LogLevel
	Transports []transport.Transport
	Loggers    map[string]LoggerConfig
}

type LoggerConfig struct {
	Level      slf4go.LogLevel
	Transports []string
}
