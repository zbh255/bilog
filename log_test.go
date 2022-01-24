package bilog

import (
	"bytes"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

func TestFeature(t *testing.T) {
	logger := NewLogger(os.Stdout, PANIC)
	logger.Info("hello world")
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
		}
	})
	b.Run("BiLogDoubleSwitchPrefix", func(b *testing.B) {
		b.ReportAllocs()
		logger := NewLogger(&TestWriter{}, PANIC)
		for i := 0; i < b.N; i++ {
			logger.Info("hello world")
			logger.Debug("hello world!")
			logger.Flush()
		}
	})
	b.Run("StdLog", func(b *testing.B) {
		b.ReportAllocs()
		logger := log.New(&TestWriter{}, "[Error] ", log.LstdFlags)
		for i := 0; i < b.N; i++ {
			logger.Print("hello world")
		}
	})
	b.Run("StdLogDoubleSwitchPrefix", func(b *testing.B) {
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

func TestSync(t *testing.T) {
	logger := NewLogger(&TestSyncWriter{}, PANIC)
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
	t.Log(time.Duration(times))
}
