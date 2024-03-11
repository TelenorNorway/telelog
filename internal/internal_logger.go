package internal

import (
	"fmt"
	"github.com/telenornorway/slf4go"
	"github.com/telenornorway/telelog/encode/logfmt"
	"github.com/telenornorway/telelog/transport"
	"runtime"
	"runtime/debug"
	"strings"
)

type LogLevel int

const (
	Inherit LogLevel = -1
	Fatal   LogLevel = 0
	Error   LogLevel = 1
	Warn    LogLevel = 2
	Info    LogLevel = 3
	Debug   LogLevel = 4
	Trace   LogLevel = 5
)

var loggers = make(map[string]*Logger)

var RootLogger = func() (logger *Logger) {
	logger = &Logger{
		name:  "",
		level: Info,
	}
	return
}()

type Logger struct {
	parent     *Logger
	name       string
	level      LogLevel
	transports []*transport.Transport
}

//goland:noinspection GoMixedReceiverTypes
func (l *Logger) SetLogLevel(level slf4go.LogLevel) {
	l.level = LogLevel(level)
}

//goland:noinspection GoMixedReceiverTypes
func (l *Logger) GetLogLevel() slf4go.LogLevel {
	if l.level == Inherit {
		return l.parent.GetLogLevel()
	}
	return slf4go.LogLevel(l.level)
}

//goland:noinspection GoMixedReceiverTypes
func (l *Logger) GetName() string {
	return l.name
}

//goland:noinspection GoMixedReceiverTypes
func (l *Logger) IsEnabled(level slf4go.LogLevel) bool {
	return l.GetLogLevel() <= level
}

func GetLogCaller() (functionPackage string, functionName string, file string, line int) {
	var pc uintptr
	var ok bool
	if pc, file, line, ok = runtime.Caller(3); !ok {
		panic("could not get caller info when logging")
	} else {
		fn := runtime.FuncForPC(pc)
		if info, ok := debug.ReadBuildInfo(); !ok {
			panic("could not get build info")
		} else {
			for _, module := range info.Deps {
				println(module.Path, module.Version)
			}
		}
		functionName = fn.Name()
		println(functionName, file, line)
		debug.PrintStack()
	}
	return
}

//goland:noinspection GoMixedReceiverTypes
func (l *Logger) Output(level slf4go.LogLevel, format string, args ...any) {
	GetLogCaller()
	fmt.Printf("[%s] (%s) "+format+"\n", append([]any{l.name, logfmt.LevelToString(level)}, args...)...)
}

//goland:noinspection GoMixedReceiverTypes
func (l Logger) String() string {
	return fmt.Sprintf("Logger[%s]", l.name)
}

func GetInternalLoggerFor(name string) (logger *Logger) {
	var initialized bool
	if initialized, logger = getLoggerByName(name); !initialized {
		initializeLogger(logger)
	}
	return logger
}

// internal-internals past this

func getLoggerByName(name string) (initialized bool, logger *Logger) {
	if logger, ok := loggers[name]; ok {
		return true, logger
	}

	logger = &Logger{
		name:  name,
		level: Inherit,
	}

	loggers[name] = logger
	return false, logger
}

func initializeLogger(logger *Logger) {
	var parentNames = getParentLoggerNames(logger.name)
	var previous = RootLogger
	var initialized bool
	var current *Logger
	for index := len(parentNames) - 1; index >= 0; index-- {
		if initialized, current = getLoggerByName(parentNames[index]); !initialized {
			current.parent = previous
		}
		previous = current
	}
	logger.parent = previous
}

func getParentLoggerNames(path string) (names []string) {
	var slash = strings.LastIndex(path, "/")
	for slash > 0 {
		path = path[:slash]
		names = append(names, path)
		slash = strings.LastIndex(path, "/")
	}
	return
}
