package bilog

type loggerConfig struct {
	tt *timeTemplate
	st *sourceTemplate
}

// 输出时间的模板
type timeTemplate struct {
	start bool
	region int
}

// 显示源代码行号的模板
type sourceTemplate struct {
	start bool
	split byte
}

