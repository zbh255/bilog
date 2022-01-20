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
	BUFFER_SIZE     = 256
	TIME_BUF_SIZE   = 64
	PREFIX_BUF_SIZE = 64
)

// 日志的级别


// Logger atomic flag 会有公平性的问题
type Logger struct {
	mu      sync.Mutex
	factory *TimeFactory
	level   int
	write   io.Writer
	timeBuf []byte
	prefix  []byte
	buf     []byte
}

func NewLogger(write io.Writer) *Logger {
	factory := &TimeFactory{}
	factory.Start()
	return &Logger{
		level:   0,
		write:   write,
		buf:     make([]byte, BUFFER_SIZE),
		prefix:  make([]byte, 0, PREFIX_BUF_SIZE),
		timeBuf: make([]byte, 0, TIME_BUF_SIZE),
		factory: factory,
	}
}

func (l *Logger) SetPrefix(s string) {
	//for !atomic.CompareAndSwapInt32(l.flas, NoWriter,Writer) {}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.resetPrefixBuf()
	l.prefix = append(l.prefix,s...)
}

func (l *Logger) fastConvert() {
	l.resetTimeBuf()
	now := l.factory.Get()
	yy, mm, dd := now.Date()
	l.timeBuf = append(l.timeBuf, fastConvertYear(yy)...)
	l.timeBuf = append(l.timeBuf, fastConvertMonth(int(mm))...)
	l.timeBuf = append(l.timeBuf, fastConvertDay(dd)...)
}

// 重置用于保存时间的缓冲区
func (l *Logger) resetTimeBuf() {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&l.timeBuf))
	h.Len = 0
	l.timeBuf = *(*[]byte)(unsafe.Pointer(h))
}

// 重置用于输出的缓冲区
func (l *Logger) resetWriteBuf() {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&l.buf))
	h.Len = 0
	l.buf = *(*[]byte)(unsafe.Pointer(h))
}

// 重置前缀的缓冲区
func (l *Logger) resetPrefixBuf()  {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&l.prefix))
	h.Len = 0
	l.prefix = *(*[]byte)(unsafe.Pointer(h))
}

func (l *Logger) Print(s string) {
	//for !atomic.CompareAndSwapInt32(l.flas, NoWriter,Writer) {}
	//defer func() {
	//	*l.flas = NoWriter
	//}()
	l.mu.Lock()
	defer l.mu.Unlock()

	l.resetWriteBuf()
	// 获取日志的时间
	l.fastConvert()
	l.buf = append(l.buf, l.timeBuf...)
	l.buf = append(l.buf, l.prefix...)
	l.buf = append(l.buf, s...)

	_, err := l.write.Write(l.buf)
	if err != nil {
		panic(err)
	}
	//n := len(s) + len(l.prefix)
	//sBytes := *(*[]byte)(unsafe.Pointer(&s))
	//bufTmp := (*reflect.SliceHeader)(unsafe.Pointer(&l.buf))
	//copy(l.buf, l.prefix)
	//// 下一次拷贝从上次结束的位置开始
	//bufTmp.Data = bufTmp.Data + uintptr(len(l.prefix))
	//copy(*(*[]byte)(unsafe.Pointer(bufTmp)), sBytes)
	//// 恢复指针
	//bufTmp.Data = bufTmp.Data - uintptr(len(l.prefix))
	//// 设置正确的长度并写入
	//bufTmp.Len = n
	//_, err := l.write.Write(*(*[]byte)(unsafe.Pointer(bufTmp)))
	//if err != nil {
	//	panic(err)
	//}
}
