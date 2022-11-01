package log

import (
	"strconv"

	"github.com/rs/zerolog"
)

// Config log config.
type Config struct {
	Lvl zerolog.Level
	// Enable specified log handlers
	StdoutLog bool

	FileLog bool
	// FileLog Absolute Pathname
	FileConfig struct {
		Filename string
	}

	// Todo 特定字段key加黑，日志记录只显示***，譬如各种token用户隐私等
	Filters []string
}

type SlsConfig struct {
	Project  string
	LogStore string
}

var handlers = make([]Handler, 0, 1)

func Init(conf *Config) {
	if conf.FileLog {
		fileHandler := NewFile(conf.FileConfig.Filename)
		handlers = append(handlers, fileHandler)
	}
	if conf.StdoutLog {
		stdoutHandler := NewStdout()
		handlers = append(handlers, stdoutHandler)
	}

	for _, h := range handlers {
		h.Level(conf.Lvl)
	}
}

func withLevel(h Handler, level zerolog.Level) LogEvent {
	switch level {
	case zerolog.DebugLevel:
		return h.Debug()
	case zerolog.InfoLevel:
		return h.Info()
	case zerolog.WarnLevel:
		return h.Warn()
	case zerolog.ErrorLevel:
		return h.Error()
	case zerolog.FatalLevel:
		return h.Fatal()
	case zerolog.PanicLevel:
		return h.Panic()
	case zerolog.Disabled:
		return nil
	default:
		panic("basic log: withLevel(): invalid level: " + strconv.Itoa(int(level)))
	}
}

func makeEvents(level zerolog.Level) *LogEvents {
	les := LogEvents{logEvents: make([]LogEvent, 0, len(handlers))}
	for _, h := range handlers {
		le := withLevel(h, level)
		les.logEvents = append(les.logEvents, le)
	}
	return &les
}

func Debug() *LogEvents {
	return makeEvents(zerolog.DebugLevel)
}

func Info() *LogEvents {
	return makeEvents(zerolog.InfoLevel)
}

func Warn() *LogEvents {
	return makeEvents(zerolog.WarnLevel)
}

func Error() *LogEvents {
	return makeEvents(zerolog.ErrorLevel)
}

func Fatal() *LogEvents {
	return makeEvents(zerolog.FatalLevel)
}

func Panic() *LogEvents {
	return makeEvents(zerolog.PanicLevel)
}

func Err(err error) *LogEvents {
	les := LogEvents{logEvents: make([]LogEvent, 0, len(handlers))}
	for _, h := range handlers {
		les.logEvents = append(les.logEvents, h.Err(err))
	}
	return &les
}

type handlerErrors []error

func (hes handlerErrors) Error() string {
	msg := ""
	for _, he := range hes {
		msg += he.Error() + "\n"
	}
	return msg
}

func Close() error {
	hes := make(handlerErrors, 0, 1)
	for _, h := range handlers {
		err := h.Close()
		if err != nil {
			hes = append(hes, err)
		}
	}
	return hes
}
