package bilog

/*
	这个包里定义了bilog中抽象的接口
*/

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

type options interface {
	apply(option *loggerConfig)
}

// 检查是否实现Logger接口
// 通过接口调用方法会降低性能，但为了约束api的原因
// 所以需要在此处检查
func init() {
	_ = Logger(&SimpleLogger{})
}
