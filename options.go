package bilog

type loggerConfig struct {
	tt timeTemplate
	st            sourceTemplate
	lowBufferSize int
	topBufferSize int
}

// 输出时间的模板
type timeTemplate struct {
	start  bool
	region int
}

// 显示源代码行号的模板
type sourceTemplate struct {
	start bool
	split byte
}

type WithFunc func(options *loggerConfig)

func (w WithFunc) apply(options *loggerConfig) {
	w(options)
}

func WithDefault() WithFunc {
	return func(options *loggerConfig) {
		options.tt.start = true
		options.st.start = false
		options.lowBufferSize = DEFAULT_LOW_BUFFER_SIZE
		options.topBufferSize = DEFAULT_TOP_BUFFER_SIZE
	}
}

func WithCaller() WithFunc {
	return func(options *loggerConfig) {
		options.st.start = true
	}
}

func WithTimes() WithFunc {
	return func(options *loggerConfig) {
		options.tt.start = true
	}
}

// WithLowBuffer 大小可已设置为 N * DEFAULT_TOP_BUFFER_SIZE
// nTopBuffer == N
func WithLowBuffer(nTopBuffer int8) WithFunc {
	return func(options *loggerConfig) {
		options.lowBufferSize = int(nTopBuffer) * DEFAULT_TOP_BUFFER_SIZE
	}
}

// WithTopBuffer 原来的pow语义不够清晰，设置0-Buffer时无法提供一个清晰的语义
// 所以将其改成直接设置Top-Buffer的大小
func WithTopBuffer(lowBufferSize int32) WithFunc {
	return func(options *loggerConfig) {
		options.topBufferSize = int(lowBufferSize)
	}
}