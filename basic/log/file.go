package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type FileHandler struct {
	logger *zerolog.Logger
	file   *os.File
}

type FileLogEvent struct {
	event *zerolog.Event
}

func NewFile(filename string) *FileHandler {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(fmt.Sprintf("Can't open log file %s for basic logger. Err: %s", filename, err))
	}
	logger := zerolog.New(f).With().Timestamp().Logger()
	return &FileHandler{
		logger: &logger,
		file: f,
	}
}

func (f *FileHandler) Level(lvl zerolog.Level) Handler {
	newLogger := f.logger.Level(lvl)
	f.logger = &newLogger
	return f
}

func (f *FileHandler) GetLevel() zerolog.Level {
	return f.logger.GetLevel()
}

func (f *FileHandler) Debug() LogEvent {
	return &FileLogEvent{event: f.logger.Debug()}
}

func (f *FileHandler) Info() LogEvent {
	return &FileLogEvent{event: f.logger.Info()}
}

func (f *FileHandler) Warn() LogEvent {
	return &FileLogEvent{event: f.logger.Warn()}
}

func (f *FileHandler) Error() LogEvent {
	return &FileLogEvent{event: f.logger.Error()}
}

func (f *FileHandler) Fatal() LogEvent {
	return &FileLogEvent{event: f.logger.Fatal()}
}

func (f *FileHandler) Panic() LogEvent {
	return &FileLogEvent{event: f.logger.Panic()}
}

func (f *FileHandler) Err(err error) LogEvent {
	return &FileLogEvent{event: f.logger.Err(err)}
}

func (f *FileHandler) Close() error {
	return f.file.Close()
}

//Events
func (fle *FileLogEvent) Msg(msg string) {
	fle.event.Msg(msg)
}

func (fle *FileLogEvent) Err(err error) LogEvent {
	fle.event.Err(err)
	return fle
}

func (fle *FileLogEvent) Str(key, val string) LogEvent {
	fle.event.Str(key, val)
	return fle
}

func (fle *FileLogEvent) Bool(key string, b bool) LogEvent {
	fle.event.Bool(key, b)
	return fle
}

func (fle *FileLogEvent) Int(key string, i int) LogEvent {
	fle.event.Int(key, i)
	return fle
}

func (fle *FileLogEvent) Int8(key string, i int8) LogEvent {
	fle.event.Int8(key, i)
	return fle
}

func (fle *FileLogEvent) Int16(key string, i int16) LogEvent {
	fle.event.Int16(key, i)
	return fle
}

func (fle *FileLogEvent) Int32(key string, i int32) LogEvent {
	fle.event.Int32(key, i)
	return fle
}

func (fle *FileLogEvent) Int64(key string, i int64) LogEvent {
	fle.event.Int64(key, i)
	return fle
}

func (fle *FileLogEvent) Uint(key string, i uint) LogEvent {
	fle.event.Uint(key, i)
	return fle
}

func (fle *FileLogEvent) Uint8(key string, i uint8) LogEvent {
	fle.event.Uint8(key, i)
	return fle
}

func (fle *FileLogEvent) Uint16(key string, i uint16) LogEvent {
	fle.event.Uint16(key, i)
	return fle
}

func (fle *FileLogEvent) Uint32(key string, i uint32) LogEvent {
	fle.event.Uint32(key, i)
	return fle
}

func (fle *FileLogEvent) Uint64(key string, i uint64) LogEvent {
	fle.event.Uint64(key, i)
	return fle
}

func (fle *FileLogEvent) Float32(key string, f float32) LogEvent {
	fle.event.Float32(key, f)
	return fle
}

func (fle *FileLogEvent) Float64(key string, f float64) LogEvent {
	fle.event.Float64(key, f)
	return fle
}

func (fle *FileLogEvent) Time(key string, t time.Time) LogEvent {
	fle.event.Time(key, t)
	return fle
}

func (fle *FileLogEvent) Dur(key string, d time.Duration) LogEvent {
	fle.event.Dur(key, d)
	return fle
}

func (fle *FileLogEvent) Interface(key string, i interface{}) LogEvent {
	fle.event.Interface(key, i)
	return fle
}

func (fle *FileLogEvent) Caller() LogEvent {
	fle.event.Caller()
	return fle
}

