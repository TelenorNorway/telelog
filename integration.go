package telelog

import (
	"github.com/telenornorway/slf4go"
	"github.com/telenornorway/telelog/internal"
)

type telelogDriver struct{}

func (t telelogDriver) Name() string {
	return "telelog"
}

func (t telelogDriver) GetLogger() slf4go.Logger {
	return &logger{internalLogger: internal.RootLogger}
}

func (t telelogDriver) MdcClear() {
	panic("implement me")
}

func (t telelogDriver) MdcPut(key, value string) {
	panic("implement me")
}

func (t telelogDriver) MdcGet(key string) (string, bool) {
	panic("implement me")
}

func (t telelogDriver) MdcRemove(key string) {
	panic("implement me")
}

func (t telelogDriver) MdcCopy() map[string]string {
	panic("implement me")
}

func Initialize(
	on func(slf4go.Driver),
	config Config,
) {
	on(telelogDriver{})
	internal.RootLogger.SetLogLevel(config.Level)
}

type logger struct {
	internalLogger *internal.Logger
}

func (l logger) Name() string {
	return l.internalLogger.GetName()
}

func (l logger) Level() slf4go.LogLevel {
	return l.internalLogger.GetLogLevel()
}

func (l logger) Trace(format string, args ...any) {
	l.internalLogger.Output(slf4go.Trace, format, args...)
}

func (l logger) TraceIf(expression bool, format string, args ...any) {
	if expression {
		l.internalLogger.Output(slf4go.Trace, format, args...)
	}
}

func (l logger) TraceUnless(expression bool, format string, args ...any) {
	if !expression {
		l.internalLogger.Output(slf4go.Trace, format, args...)
	}
}

func (l logger) IsTraceEnabled() bool {
	return l.internalLogger.IsEnabled(slf4go.Trace)
}

func (l logger) Debug(format string, args ...any) {
	l.internalLogger.Output(slf4go.Debug, format, args...)
}

func (l logger) DebugIf(expression bool, format string, args ...any) {
	if expression {
		l.internalLogger.Output(slf4go.Debug, format, args...)
	}
}

func (l logger) DebugUnless(expression bool, format string, args ...any) {
	if !expression {
		l.internalLogger.Output(slf4go.Debug, format, args...)
	}
}

func (l logger) IsDebugEnabled() bool {
	return l.internalLogger.IsEnabled(slf4go.Debug)
}

func (l logger) Info(format string, args ...any) {
	l.internalLogger.Output(slf4go.Info, format, args...)
}

func (l logger) InfoIf(expression bool, format string, args ...any) {
	if expression {
		l.internalLogger.Output(slf4go.Info, format, args...)
	}
}

func (l logger) InfoUnless(expression bool, format string, args ...any) {
	if !expression {
		l.internalLogger.Output(slf4go.Info, format, args...)
	}
}

func (l logger) IsInfoEnabled() bool {
	return l.internalLogger.IsEnabled(slf4go.Info)
}

func (l logger) Warn(format string, args ...any) {
	l.internalLogger.Output(slf4go.Warn, format, args...)
}

func (l logger) WarnIf(expression bool, format string, args ...any) {
	if expression {
		l.internalLogger.Output(slf4go.Warn, format, args...)
	}
}

func (l logger) WarnUnless(expression bool, format string, args ...any) {
	if !expression {
		l.internalLogger.Output(slf4go.Warn, format, args...)
	}
}

func (l logger) IsWarnEnabled() bool {
	return l.internalLogger.IsEnabled(slf4go.Warn)
}

func (l logger) Error(format string, args ...any) {
	l.internalLogger.Output(slf4go.Error, format, args...)
}

func (l logger) ErrorIf(expression bool, format string, args ...any) {
	if expression {
		l.internalLogger.Output(slf4go.Error, format, args...)
	}
}

func (l logger) ErrorUnless(expression bool, format string, args ...any) {
	if !expression {
		l.internalLogger.Output(slf4go.Error, format, args...)
	}
}

func (l logger) IsErrorEnabled() bool {
	return l.internalLogger.IsEnabled(slf4go.Error)
}

func (l logger) Fatal(format string, args ...any) {
	l.internalLogger.Output(slf4go.Fatal, format, args...)
}

func (l logger) FatalIf(expression bool, format string, args ...any) {
	if expression {
		l.internalLogger.Output(slf4go.Fatal, format, args...)
	}
}

func (l logger) FatalUnless(expression bool, format string, args ...any) {
	if !expression {
		l.internalLogger.Output(slf4go.Fatal, format, args...)
	}
}

func (l logger) IsFatalEnabled() bool {
	return l.internalLogger.IsEnabled(slf4go.Fatal)
}
