package main

import (
	"github.com/telenornorway/slf4go"
	"github.com/telenornorway/telelog"
	"github.com/telenornorway/telelog/encode/logfmt"
	"github.com/telenornorway/telelog/transport"
	"github.com/telenornorway/telelog/transport/console"
)

func main() {
	telelog.Initialize(slf4go.UseDriver, telelog.Config{
		Level:      slf4go.Info,
		Transports: []transport.Transport{console.NewConsoleTransport(logfmt.New())},
		Loggers:    map[string]telelog.LoggerConfig{},
	})

	var logger = slf4go.GetLogger()

	logger.Info("Hello, World!")
}
