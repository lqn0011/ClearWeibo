package log

import "time"

type (
	LogEvent interface {
		Msg(msg string)

		Err(err error) LogEvent

		Str(key, val string) LogEvent

		Bool(key string, b bool) LogEvent

		Int(key string, i int) LogEvent
		Int8(key string, i int8) LogEvent
		Int16(key string, i int16) LogEvent
		Int32(key string, i int32) LogEvent
		Int64(key string, i int64) LogEvent

		Uint(key string, i uint) LogEvent
		Uint8(key string, i uint8) LogEvent
		Uint16(key string, i uint16) LogEvent
		Uint32(key string, i uint32) LogEvent
		Uint64(key string, i uint64) LogEvent

		Float32(key string, f float32) LogEvent
		Float64(key string, f float64) LogEvent

		Time(key string, t time.Time) LogEvent
		Dur(key string, d time.Duration) LogEvent

		Interface(key string, i interface{}) LogEvent
		Caller() LogEvent
	}

	LogEvents struct {
		logEvents []LogEvent
	}
)

var (
	_ LogEvent = &FileLogEvent{}
	_ LogEvent = &StdoutLogEvent{}
)

func (les *LogEvents) Msg(msg string) {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Msg(msg)
	}
}

func (les *LogEvents) Err(err error) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Err(err)
	}
	return les
}

func (les *LogEvents) Str(key, val string) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Str(key, val)
	}
	return les
}

func (les *LogEvents) Bool(key string, b bool) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Bool(key, b)
	}
	return les
}

func (les *LogEvents) Int(key string, i int) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Int(key, i)
	}
	return les
}

func (les *LogEvents) Int8(key string, i int8) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Int8(key, i)
	}
	return les
}

func (les *LogEvents) Int16(key string, i int16) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Int16(key, i)
	}
	return les
}

func (les *LogEvents) Int32(key string, i int32) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Int32(key, i)
	}
	return les
}

func (les *LogEvents) Int64(key string, i int64) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Int64(key, i)
	}
	return les
}

func (les *LogEvents) Uint(key string, i uint) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Uint(key, i)
	}
	return les
}

func (les *LogEvents) Uint8(key string, i uint8) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Uint8(key, i)
	}
	return les
}

func (les *LogEvents) Uint16(key string, i uint16) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Uint16(key, i)
	}
	return les
}

func (les *LogEvents) Uint32(key string, i uint32) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Uint32(key, i)
	}
	return les
}

func (les *LogEvents) Uint64(key string, i uint64) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Uint64(key, i)
	}
	return les
}

func (les *LogEvents) Float32(key string, f float32) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Float32(key, f)
	}
	return les
}

func (les *LogEvents) Float64(key string, f float64) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Float64(key, f)
	}
	return les
}

func (les *LogEvents) Time(key string, t time.Time) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Time(key, t)
	}
	return les
}

func (les *LogEvents) Dur(key string, d time.Duration) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Dur(key, d)
	}
	return les
}

func (les *LogEvents) Interface(key string, i interface{}) *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Interface(key, i)
	}
	return les
}

func (les *LogEvents) Caller() *LogEvents {
	for _, l := range les.logEvents {
		if l == nil {
			continue
		}
		l.Caller()
	}
	return les
}


