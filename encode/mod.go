package encode

import "github.com/telenornorway/slf4go"

type Encoder interface {
	Encode(payload *slf4go.LogPayload) []byte
}
