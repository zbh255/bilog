package bilog

import (
	"runtime"
	"sync"
	_ "unsafe"
)

const (
	defaultCallDepth = 3
	defaultFile      = "???"
)

var (
	cache           = make(map[uintptr]runtime.Frame, 16)
	concurrentCache sync.Map
)

// Caller 输出调用函数的文件名和行号
func Caller(callDepth int) (file string, line int) {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = defaultFile
	} else {
		file = cutFileName(file)
	}
	return
}

func CallerOfCache(skip int) (file string, line int) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return
	}
	if frame, ok := cache[rpc[0]]; ok {
		file = frame.File
		line = frame.Line
	} else {
		frame, _ = runtime.CallersFrames(rpc).Next()
		cache[rpc[0]] = frame
		file = frame.File
		line = frame.Line
	}
	file = cutFileName(file)
	return
}

func CallerOfConcurrentCache(skip int) (file string, line int) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return
	}
	frameI, ok := concurrentCache.Load(rpc[0])
	if ok {
		file = frameI.(runtime.Frame).File
		line = frameI.(runtime.Frame).Line
	} else {
		frame, _ := runtime.CallersFrames(rpc).Next()
		concurrentCache.Store(rpc[0], frame)
		file = frame.File
		line = frame.Line
	}
	file = cutFileName(file)
	return
}

func cutFileName(file string) string {
	switch file {
	case "", defaultFile:
		return defaultFile
	}
	var split int
	var match int
	for i := len(file) - 1; i >= 0; i-- {
		if match == 2 {
			split = i + 2
			break
		}
		switch file[i] {
		case '/', '\\':
			match++
		}
	}
	return file[split:]
}
