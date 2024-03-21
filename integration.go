package telelog

import (
	"fmt"

	"github.com/telenornorway/mdc"
	"github.com/telenornorway/slf4go"
	"github.com/telenornorway/telelog/encode/logfmt"
	"github.com/telenornorway/telelog/internal"
)

type telelogDriver struct{}

func (t telelogDriver) Name() string { return "telelog" }

func (t telelogDriver) GetLevelForLoggerWithName(name string) slf4go.LogLevel {
	return internal.GetInternalLoggerFor(name).GetLogLevel()
}

func (t telelogDriver) Write(payload slf4go.LogPayload) {
	fmt.Printf("%s", string(logfmt.New().Encode(&payload)))
}

func (t telelogDriver) MdcPut(key, value string) {
	mdc.Put(key, value)
}

func (t telelogDriver) MdcGet(key string) (value string, exists bool) {
	exists, value = mdc.Get(key)
	return
}

func (t telelogDriver) MdcRemove(key string) {
	mdc.Remove(key)
}

func (t telelogDriver) MdcClear() {
	mdc.Clear()
}

func (t telelogDriver) MdcCopy() map[string]string {
	return mdc.Copy()
}

func Initialize(
	on func(slf4go.Driver),
	config Config,
) {
	on(telelogDriver{})
	internal.RootLogger.SetLogLevel(config.Level)
}
