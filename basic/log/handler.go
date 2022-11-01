package log

import "github.com/rs/zerolog"

type Handler interface {
	Level(lvl zerolog.Level) Handler
	GetLevel() zerolog.Level

	Debug() LogEvent
	Info() LogEvent
	Warn() LogEvent
	Error() LogEvent
	Fatal() LogEvent
	Panic() LogEvent

	Err(err error) LogEvent

	Close() error
}

var(
	_ Handler = &FileHandler{}
	_ Handler = &StdoutHandler{}
)

