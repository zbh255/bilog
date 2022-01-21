package bilog

// 日志的级别
type level int

const (
	INFO level = iota
	DEBUG
	TRACE
	ERROR
	PANIC
)

type Logger interface {
	Level() int
	Info(s string)
	Debug(s string)
	Trace(s string)
	ErrorFromErr(e error)
	ErrorFromString(s string)
	PanicFromErr(e error)
	PanicFromString(s string)
	Flush()
}
