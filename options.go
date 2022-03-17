package bilog

type loggerConfig struct {
	tt timeTemplate
	st sourceTemplate
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
