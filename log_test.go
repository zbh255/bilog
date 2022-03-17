package bilog

import (
	"bytes"
	"errors"
	"log"
	"sync"
	"testing"
	"time"
)

// 测试等级不同的日志是否写入的写入器
type TestLevelWriter struct {
	setLevel  level
	liveLevel level
}

func (t *TestLevelWriter) Check(liveLevel level) bool {
	return t.setLevel >= liveLevel
}

func (t *TestLevelWriter) SetLiveLevel(liveLevel level) {
	t.liveLevel = liveLevel
}

func (t *TestLevelWriter) Write(p []byte) (n int, err error) {
	if !t.Check(t.liveLevel) && len(p) > 0 {
		panic("setLevel less than liveLevel")
	}
	return len(p), nil
}

func TestFeature(t *testing.T) {
	test := &TestLevelWriter{
		setLevel: TRACE,
	}
	logger := NewLogger(test, TRACE)
	// info log
	test.SetLiveLevel(INFO)
	logger.Info("hello world")
	logger.Flush()
	// debug log
	test.SetLiveLevel(DEBUG)
	logger.Debug("hello world")
	logger.Flush()
	// trace log
	test.SetLiveLevel(TRACE)
	logger.Trace("hello world")
	logger.Flush()
	// error log
	test.SetLiveLevel(ERROR)
	logger.ErrorFromErr(errors.New("my is error"))
	logger.Flush()
	// panic log
	test.SetLiveLevel(PANIC)
	logger.PanicFromString("my is panic")
	logger.Flush()
}

type TestWriter struct {
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// 测试使用的同步写入器
type TestSyncWriter struct {
	buf []byte
}

func (t *TestSyncWriter) Write(p []byte) (n int, err error) {
	if t.buf == nil {
		t.buf = make([]byte, len(p))
		copy(t.buf, p)
	}
	if !bytes.Equal(t.buf, p) {
		panic("write data is not equal topBuf")
	}
	return len(p), nil
}

func BenchmarkLogger(b *testing.B) {
	b.Run("BiLog", func(b *testing.B) {
		b.ReportAllocs()
		logger := NewLogger(&TestWriter{}, PANIC)
		for i := 0; i < b.N; i++ {
			logger.Debug("hello world")
			logger.Flush()
		}
	})
	b.Run("BiLogDouble", func(b *testing.B) {
		b.ReportAllocs()
		logger := NewLogger(&TestWriter{}, PANIC)
		for i := 0; i < b.N; i++ {
			logger.Info("hello world")
			logger.Debug("hello world!")
			logger.Flush()
		}
	})
	b.Run("BilogCallerAndTime", func(b *testing.B) {
		b.ReportAllocs()
		logger := NewLogger(&TestWriter{}, PANIC,
			WithTimes(),
			WithCaller(),
		)
		for i := 0; i < b.N; i++ {
			logger.Info("hello world")
		}
	})
	b.Run("StdLogCallerAndTime", func(b *testing.B) {
		b.ReportAllocs()
		logger := log.New(&TestWriter{}, "[PANIC]", log.Llongfile|log.Ltime)
		for i := 0; i < b.N; i++ {
			logger.Println("hello world!")
		}
	})
	b.Run("StdLog", func(b *testing.B) {
		b.ReportAllocs()
		logger := log.New(&TestWriter{}, "[Error] ", log.LstdFlags)
		for i := 0; i < b.N; i++ {
			logger.Print("hello world")
		}
	})
	b.Run("StdLogDouble", func(b *testing.B) {
		b.ReportAllocs()
		logger := log.New(&TestWriter{}, "[Error] ", log.LstdFlags)
		for i := 0; i < b.N; i++ {
			logger.SetPrefix("[INFO] ")
			logger.Println("hello world")
			logger.SetPrefix("[DEBUG] ")
			logger.Println("hello world")
		}
	})
}

func BenchmarkCallerLogger(b *testing.B) {
	b.Run("StdLog", func(b *testing.B) {
		b.ReportAllocs()
		logger := log.New(&TestWriter{}, "[Error]", log.Llongfile|log.Ltime)
		for i := 0; i < b.N; i++ {
			logger.Println("hello world!")
		}
	})
}

func TestSync(t *testing.T) {
	logger := NewLogger(&TestWriter{}, PANIC)
	// goroutine等待的最长时间
	var times int64
	// 保护等待时间的互斥锁
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func() {
			defer wg.Done()
			unixSt := time.Now().UnixNano()
			logger.Info("hello world")
			unixEnd := time.Now().UnixNano()
			mu.Lock()
			if unixEnd-unixSt > times {
				times = unixEnd - unixSt
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	// 如果goroutine等待锁的时间太长则测试失败
	if time.Duration(times) > time.Second {
		t.Error("goroutine wait time out")
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/zbh255/bilog
cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
BenchmarkInterfaceCall
BenchmarkInterfaceCall/Interfaces
BenchmarkInterfaceCall/Interfaces-8         	31814370	        40.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfaceCall/NoInterfaces
BenchmarkInterfaceCall/NoInterfaces-8       	36872288	        33.63 ns/op	       0 B/op	       0 allocs/op
PASS
*/
// 测试发现，在bilog中通过接口调用会带来10%的性能损失，这一部分可以优化
func BenchmarkInterfaceCall(b *testing.B) {
	b.Run("Interfaces", func(b *testing.B) {
		b.ReportAllocs()
		logger := Logger(NewLogger(&TestWriter{}, PANIC))
		for i := 0; i < b.N; i++ {
			logger.Info("hello world!")
		}
	})
	b.Run("NoInterfaces", func(b *testing.B) {
		b.ReportAllocs()
		logger := NewLogger(&TestWriter{}, PANIC)
		for i := 0; i < b.N; i++ {
			logger.Info("hello world!")
		}
	})
}
