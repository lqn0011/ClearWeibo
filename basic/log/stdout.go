package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type StdoutHandler struct {
	logger *zerolog.Logger
}

type StdoutLogEvent struct {
	event *zerolog.Event
}

func NewStdout() *StdoutHandler {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	return &StdoutHandler{
		logger: &logger,
	}
}

func (f *StdoutHandler) Level(lvl zerolog.Level) Handler {
	newLogger := f.logger.Level(lvl)
	f.logger = &newLogger
	return f
}

func (f *StdoutHandler) GetLevel() zerolog.Level {
	return f.logger.GetLevel()
}

func (f *StdoutHandler) Debug() LogEvent {
	return &StdoutLogEvent{event: f.logger.Debug()}
}

func (f *StdoutHandler) Info() LogEvent {
	return &StdoutLogEvent{event: f.logger.Info()}
}

func (f *StdoutHandler) Warn() LogEvent {
	return &StdoutLogEvent{event: f.logger.Warn()}
}

func (f *StdoutHandler) Error() LogEvent {
	return &StdoutLogEvent{event: f.logger.Error()}
}

func (f *StdoutHandler) Fatal() LogEvent {
	return &StdoutLogEvent{event: f.logger.Fatal()}
}

func (f *StdoutHandler) Panic() LogEvent {
	return &StdoutLogEvent{event: f.logger.Panic()}
}

func (f *StdoutHandler) Err(err error) LogEvent {
	return &StdoutLogEvent{event: f.logger.Err(err)}
}

func (f *StdoutHandler) Close() error {
	return nil
}

//Events
func (fle *StdoutLogEvent) Msg(msg string) {
	fle.event.Msg(msg)
}

func (fle *StdoutLogEvent) Err(err error) LogEvent {
	fle.event.Err(err)
	return fle
}

func (fle *StdoutLogEvent) Str(key, val string) LogEvent {
	fle.event.Str(key, val)
	return fle
}

func (fle *StdoutLogEvent) Bool(key string, b bool) LogEvent {
	fle.event.Bool(key, b)
	return fle
}

func (fle *StdoutLogEvent) Int(key string, i int) LogEvent {
	fle.event.Int(key, i)
	return fle
}

func (fle *StdoutLogEvent) Int8(key string, i int8) LogEvent {
	fle.event.Int8(key, i)
	return fle
}

func (fle *StdoutLogEvent) Int16(key string, i int16) LogEvent {
	fle.event.Int16(key, i)
	return fle
}

func (fle *StdoutLogEvent) Int32(key string, i int32) LogEvent {
	fle.event.Int32(key, i)
	return fle
}

func (fle *StdoutLogEvent) Int64(key string, i int64) LogEvent {
	fle.event.Int64(key, i)
	return fle
}

func (fle *StdoutLogEvent) Uint(key string, i uint) LogEvent {
	fle.event.Uint(key, i)
	return fle
}

func (fle *StdoutLogEvent) Uint8(key string, i uint8) LogEvent {
	fle.event.Uint8(key, i)
	return fle
}

func (fle *StdoutLogEvent) Uint16(key string, i uint16) LogEvent {
	fle.event.Uint16(key, i)
	return fle
}

func (fle *StdoutLogEvent) Uint32(key string, i uint32) LogEvent {
	fle.event.Uint32(key, i)
	return fle
}

func (fle *StdoutLogEvent) Uint64(key string, i uint64) LogEvent {
	fle.event.Uint64(key, i)
	return fle
}

func (fle *StdoutLogEvent) Float32(key string, f float32) LogEvent {
	fle.event.Float32(key, f)
	return fle
}

func (fle *StdoutLogEvent) Float64(key string, f float64) LogEvent {
	fle.event.Float64(key, f)
	return fle
}

func (fle *StdoutLogEvent) Time(key string, t time.Time) LogEvent {
	fle.event.Time(key, t)
	return fle
}

func (fle *StdoutLogEvent) Dur(key string, d time.Duration) LogEvent {
	fle.event.Dur(key, d)
	return fle
}

func (fle *StdoutLogEvent) Interface(key string, i interface{}) LogEvent {
	fle.event.Interface(key, i)
	return fle
}

func (fle *StdoutLogEvent) Caller() LogEvent {
	fle.event.Caller()
	return fle
}

