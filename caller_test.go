package bilog

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkCaller(b *testing.B) {
	b.Run("Default", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = Caller(1)
		}
	})
	b.Run("CallerCached", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = CallerOfCache(1)
		}
	})
}

func TestCaller(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			caller()
			caller()
		}()
	}
	wg.Wait()
}

func caller() {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("pc=%d,file=%s,line=%d\n", pc, file, line)
}
