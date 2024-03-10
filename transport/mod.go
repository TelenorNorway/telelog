package transport

import (
	"github.com/telenornorway/telelog/encode"
	"io"
	"sync"
)

type Transport interface {
	GetEncoder() encode.Encoder
	Write(payload []byte)
}

type anonymousTransport struct {
	encoder encode.Encoder
	stream  io.Writer
	mutex   *sync.Mutex
}

func (a anonymousTransport) GetEncoder() encode.Encoder {
	return a.encoder
}

func (a anonymousTransport) Write(payload []byte) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.stream.Write(payload)
}

func NewTransport(encoder encode.Encoder, stream io.Writer) Transport {
	return anonymousTransport{
		encoder: encoder,
		stream:  stream,
		mutex:   &sync.Mutex{},
	}
}
