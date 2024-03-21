package console

import (
	"os"

	"github.com/telenornorway/telelog/encode"
	"github.com/telenornorway/telelog/transport"
)

func NewConsoleTransport(encoder encode.Encoder) transport.Transport {
	return transport.NewTransport(encoder, os.Stdout)
}

func NewConsoleErrorTransport(encoder encode.Encoder) transport.Transport {
	return transport.NewTransport(encoder, os.Stderr)
}
