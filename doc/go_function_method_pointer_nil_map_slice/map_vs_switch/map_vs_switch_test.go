package map_vs_switch

import "testing"

type LogLevel int32

const (
	CRITICAL LogLevel = iota - 1
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
	TRACE
)

func (l LogLevel) StringSwitch() string {
	switch l {
	case CRITICAL:
		return "CRITICAL"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	default:
		panic("Unhandled loglevel")
	}
}

var lmap = map[LogLevel]string{
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	WARNING:  "WARNING",
	NOTICE:   "NOTICE",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
	TRACE:    "TRACE",
}

func (l LogLevel) StringMap() string {
	v, ok := lmap[l]
	if !ok {
		panic("Unhandled loglevel")
	}
	return v
}

var levels = []LogLevel{
	TRACE,
	DEBUG,
	INFO,
	NOTICE,
	WARNING,
	ERROR,
	CRITICAL,
}

func BenchmarkStringSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, l := range levels {
			l.StringSwitch()
		}
	}
}

func BenchmarkStringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, l := range levels {
			l.StringMap()
		}
	}
}

/*
BenchmarkStringSwitch-8	50000000	        20.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkStringMap-8	20000000	        62.0 ns/op	       0 B/op	       0 allocs/op
*/
