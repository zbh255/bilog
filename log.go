package bilog

import (
	"io"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"
)

// 测试版本

/*
	bilog 高性能思路：复用string参数的内存，copy但不重新分配
	prefix的内存也复用，拷贝但不分配，reset使用unsafe修改len
*/

const (
	DEFAULT_TOP_BUFFER_SIZE = 256
	DEFAULT_LOW_BUFFER_SIZE = DEFAULT_TOP_BUFFER_SIZE * 6
	TIME_BUF_SIZE           = 64
	CALLER_BUF_SIZE = 64
)

// 缓存特殊的字符
var (
	cacheEntry byte = '\n'
	cacheSplit byte = ':'
	cacheSpace byte = ' '
)

// SimpleLogger atomic flag 会有公平性的问题
type SimpleLogger struct {
	mu sync.Mutex
	// 缓存和生产date数据的工厂
	factory *TimeFactory
	// 设置的输出等级
	level level
	// config
	confObj *loggerConfig
	// 缓存level对应的string，避免频繁拷贝和分配
	// 必须按照doc.go中定义的日志等级来初始化
	levelCache []string
	write      io.Writer
	// 缓存格式化后的时间
	timeBuf []byte
	// 缓存的纳秒级时间戳
	timeStamp int64
	// 虽然runtime.Caller提供的file string已经逃逸到堆中，不用多次一举去拷贝
	// 该buffer是主要为了line而提供的
	callerBuf []byte
	//TODO:计划删除
	//prefix string
	// 顶层缓冲区
	topBuf []byte
	// 底层缓冲区
	lowBuf []byte
}

func NewLogger(write io.Writer, l level, options ...options) *SimpleLogger {
	var cf loggerConfig
	if options == nil || len(options) == 0 {
		WithDefault().apply(&cf)
	} else {
		for _, option := range options {
			option.apply(&cf)
		}
	}

	var factory *TimeFactory
	if cf.tt.start {
		factory = timeFactory
		factory.Start()
	}
	return &SimpleLogger{
		level: l,
		write: write,
		levelCache: []string{
			"[INFO] ", "[DEBUG] ", "[TRACE] ", "[ERROR] ",
			"[PANIC] ",
		},
		confObj:   &cf,
		callerBuf: make([]byte, CALLER_BUF_SIZE),
		topBuf:    make([]byte, cf.topBufferSize),
		lowBuf:    make([]byte, 0, cf.lowBufferSize),
		timeBuf:   make([]byte, 0, TIME_BUF_SIZE),
		factory:   factory,
	}
}

// SetPrefix 复用string的内存来避免memove
//func (l *SimpleLogger) SetPrefix(s string) {
//	//for !atomic.CompareAndSwapInt32(l.flas, NoWriter,Writer) {}
//	l.mu.Lock()
//	defer l.mu.Unlock()
//	l.prefix = s
//}

// TODO优化转换速度
func (l *SimpleLogger) fastConvert() {
	// 比较新老时间戳，如果还在有效时间之内则不更新timeBuf里的内容，减少memmove次数
	timeStamp := l.factory.TimeStamp()
	if !(timeStamp-l.timeStamp > int64(time.Millisecond*10)) && len(l.timeBuf) > 0 {
		l.timeStamp = timeStamp
		return
	} else {
		l.timeStamp = timeStamp
		l.resetTimeBuf()
		date := l.factory.Get()

		l.timeBuf = append(l.timeBuf, date...)
	}
}

// 重置用于保存时间的缓冲区
func (l *SimpleLogger) resetTimeBuf() {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&l.timeBuf))
	h.Len = 0
	l.timeBuf = *(*[]byte)(unsafe.Pointer(h))
}

// 写入到低层缓冲器，该函数会有一些检查
func (l *SimpleLogger) writeLowBuf() {
	if len(l.lowBuf)+len(l.topBuf) > cap(l.lowBuf) {
		l.flushLowBuf()
	}
	l.lowBuf = append(l.lowBuf, l.topBuf...)
}

// 将lowBuf中的数据全部写入到writer中并reset
func (l *SimpleLogger) flushLowBuf() {
	n, err := l.write.Write(l.lowBuf)
	if err != nil {
		panic(err)
	}
	if n != len(l.lowBuf) {
		panic("write byte not equal")
	}
	l.resetLowBuf()
}

// 重置用于输出的缓冲区
func (l *SimpleLogger) resetTopBuf() {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&l.topBuf))
	h.Len = 0
	l.topBuf = *(*[]byte)(unsafe.Pointer(h))
}

// 重置lowBuf
func (l *SimpleLogger) resetLowBuf() {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&l.lowBuf))
	h.Len = 0
	l.lowBuf = *(*[]byte)(unsafe.Pointer(h))
}

func (l *SimpleLogger) printTime(level level) {
	l.resetTopBuf()
	// 获取日志的时间
	if l.confObj.tt.start {
		l.fastConvert()
	}

	l.topBuf = append(l.topBuf, l.levelCache[level]...)
	l.topBuf = append(l.topBuf, l.timeBuf...)
}

// 在最顶层调用以提升性能
func (l *SimpleLogger) printCaller() {
	// reset
	l.callerBuf = l.callerBuf[:0]

	file, line := CallerOfConcurrentCache(defaultCallDepth)
	l.callerBuf = append(l.callerBuf, file...)
	l.callerBuf = append(l.callerBuf, cacheSplit)
	l.callerBuf = append(l.callerBuf, strconv.Itoa(line)...)
}

func (l *SimpleLogger) print(s string, level level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.printTime(level)
	l.topBuf = append(l.topBuf, s...)

	l.writeLowBuf()
}

func (l *SimpleLogger) println(s string, level level) {

	l.resetTopBuf()
	l.printTime(level)
	l.topBuf = append(l.topBuf, l.callerBuf...)
	l.topBuf = append(l.topBuf, cacheSpace)
	l.topBuf = append(l.topBuf, s...)
	l.topBuf = append(l.topBuf, cacheEntry)

	l.writeLowBuf()
}

// 验证输出的日志级别是否小于等于预设的级别
func (l *SimpleLogger) checkLevel(level level) bool {
	return level <= l.level
}

func (l *SimpleLogger) Level() int {
	return int(l.level)
}

func (l *SimpleLogger) Info(s string) {
	if !l.checkLevel(INFO) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(s, INFO)
}

func (l *SimpleLogger) Debug(s string) {
	if !l.checkLevel(DEBUG) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(s, DEBUG)
}

func (l *SimpleLogger) Trace(s string) {
	if !l.checkLevel(TRACE) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(s, TRACE)
}

//TODO: 优雅地处理error
func (l *SimpleLogger) ErrorFromErr(e error) {
	if !l.checkLevel(ERROR) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(e.Error(), ERROR)
}

func (l *SimpleLogger) ErrorFromString(s string) {
	if !l.checkLevel(ERROR) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(s, ERROR)
}

func (l *SimpleLogger) PanicFromErr(e error) {
	if !l.checkLevel(PANIC) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(e.Error(), ERROR)
	panic(e)
}

func (l *SimpleLogger) PanicFromString(s string) {
	if !l.checkLevel(PANIC) {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.confObj.st.start {
		l.printCaller()
	}
	l.println(s, ERROR)
	panic(s)
}

func (l *SimpleLogger) Flush() {
	l.flushLowBuf()
}
