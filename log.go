package bilog

import (
	"io"
	"reflect"
	"sync"
	"unsafe"
)

// 测试版本

/*
	bilog 高性能思路：复用string参数的内存，copy但不重新分配
	prefix的内存也复用，拷贝但不分配，reset使用unsafe修改len
*/

const (
	TOP_BUFFER_SIZE = 256
	LOW_BUFFER_SIZE = TOP_BUFFER_SIZE * 6
	TIME_BUF_SIZE   = 64
	PREFIX_BUF_SIZE = 64
)

// 缓存特殊的字符
var (
	cacheEntry byte = '\n'
)

// SimpleLogger atomic flag 会有公平性的问题
type SimpleLogger struct {
	mu sync.Mutex
	// 缓存和生产date数据的工厂
	factory *TimeFactory
	// 设置的输出等级
	level level
	// 缓存level对应的string，避免频繁拷贝和分配
	// 必须按照doc.go中定义的日志等级来初始化
	levelCache []string
	write      io.Writer
	// 缓存格式化后的时间
	timeBuf []byte
	//TODO:计划删除
	//prefix string
	// 顶层缓冲区
	topBuf []byte
	// 底层缓冲区
	lowBuf []byte
}

func NewLogger(write io.Writer, l level) Logger {
	factory := NewTimeFactory()
	factory.Start()
	return &SimpleLogger{
		level: l,
		write: write,
		levelCache: []string{
			"[INFO] ", "[DEBUG] ", "[TRACE] ", "[ERROR] ",
			"[PANIC] ",
		},
		topBuf:  make([]byte, TOP_BUFFER_SIZE),
		lowBuf:  make([]byte, 0, LOW_BUFFER_SIZE),
		timeBuf: make([]byte, 0, TIME_BUF_SIZE),
		factory: factory,
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
	l.resetTimeBuf()
	date := l.factory.Get()

	l.timeBuf = append(l.timeBuf, date...)
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
	_, err := l.write.Write(l.lowBuf)
	if err != nil {
		panic(err)
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

func (l *SimpleLogger) Print(s string, level level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.resetTopBuf()
	// 获取日志的时间
	l.fastConvert()
	l.topBuf = append(l.topBuf, l.levelCache[level]...)
	l.topBuf = append(l.topBuf, l.timeBuf...)
	l.topBuf = append(l.topBuf, s...)

	l.writeLowBuf()
}

func (l *SimpleLogger) Println(s string, level level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.resetTopBuf()
	// 获取日志的时间
	l.fastConvert()
	l.topBuf = append(l.topBuf, l.levelCache[level]...)
	l.topBuf = append(l.topBuf, l.timeBuf...)
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
	l.Println(s, INFO)
}

func (l *SimpleLogger) Debug(s string) {
	if !l.checkLevel(DEBUG) {
		return
	}
	l.Println(s, DEBUG)
}

func (l *SimpleLogger) Trace(s string) {
	if !l.checkLevel(TRACE) {
		return
	}
	l.Println(s, TRACE)
}

//TODO: 优雅地处理error
func (l *SimpleLogger) ErrorFromErr(e error) {
	if !l.checkLevel(ERROR) {
		return
	}
	l.Println(e.Error(), ERROR)
}

func (l *SimpleLogger) ErrorFromString(s string) {
	if !l.checkLevel(ERROR) {
		return
	}
	l.Println(s, ERROR)
}

func (l *SimpleLogger) PanicFromErr(e error) {
	if !l.checkLevel(PANIC) {
		return
	}
	l.Println(e.Error(), ERROR)
	panic(e)
}

func (l *SimpleLogger) PanicFromString(s string) {
	if !l.checkLevel(PANIC) {
		return
	}
	l.Println(s, ERROR)
	panic(s)
}

func (l *SimpleLogger) Flush() {
	l.flushLowBuf()
}
