package bilog

import (
	"runtime"
	_ "unsafe"
)

const (
	defaultCallDepth = 3
	defaultFile = "???"
)


// Caller 输出调用函数的文件名和行号
func Caller(callDepth int) (file string, line int) {
	_,file,line,ok := runtime.Caller(callDepth)
	if !ok {
		file = defaultFile
	}
	return
}
