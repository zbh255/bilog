package benchmark

import (
	"fmt"
	"github.com/zbh255/bilog"
	"log"
	"sync"
	"testing"
)

type TestWriter struct {

}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	return len(p),nil
}

func BenchmarkNoConcurrent(b *testing.B) {
	initN := 1000
	bLogger := bilog.NewLogger(&TestWriter{},bilog.PANIC,bilog.WithDefault())
	stdLogger := log.New(&TestWriter{},"[Error]",log.LstdFlags)
	// bilog buffer
	for i := 1; i <= 100; i *= 10 {
		b.Run(fmt.Sprintf("Bilog-Buffer-%d",initN * i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				for k := 0; k < initN * i; k++ {
					bLogger.ErrorFromString("hello world!")
				}
			}
		})
	}
	// bilog no-buffer
	bLogger = bilog.NewLogger(&TestWriter{},bilog.PANIC,bilog.WithTopBuffer(0),bilog.WithLowBuffer(0))
	for i := 1; i <= 1000; i *= 10 {
		b.Run(fmt.Sprintf("Bilog-NoBuffer-%d",initN * i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				for k := 0; k < initN * i; k++ {
					bLogger.ErrorFromString("hello world!")
				}
			}
		})
	}
	// stdlog no-buffer
	for i := 1; i <= 1000; i *= 10 {
		b.Run(fmt.Sprintf("Stdlog-NoBuffer-%d",initN * i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				for k := 0; k < initN * i; k++ {
					stdLogger.Println("hello world!")
				}
			}
		})
	}
}

func BenchmarkConcurrent(b *testing.B) {
	initN := 1000
	bLogger := bilog.NewLogger(&TestWriter{},bilog.PANIC,bilog.WithDefault())
	stdLogger := log.New(&TestWriter{},"[Error]",log.LstdFlags)
	// bilog buffer
	for i := 1; i <= 100; i *= 10 {
		b.Run(fmt.Sprintf("Bilog-Buffer-%d",initN * i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				var wg sync.WaitGroup
				wg.Add(initN * i)
				for k := 0; k < initN * i; k++ {
					go func() {
						bLogger.ErrorFromString("hello world!")
						wg.Done()
					}()
				}
				wg.Wait()
			}
		})
	}
	// bilog no-buffer
	bLogger = bilog.NewLogger(&TestWriter{},bilog.PANIC,bilog.WithTopBuffer(0),bilog.WithLowBuffer(0))
	for i := 1; i <= 1000; i *= 10 {
		b.Run(fmt.Sprintf("Bilog-NoBuffer-%d",initN * i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				var wg sync.WaitGroup
				wg.Add(initN * i)
				for k := 0; k < initN * i; k++ {
					go func() {
						bLogger.ErrorFromString("hello world!")
						wg.Done()
					}()
				}
				wg.Wait()
			}
		})
	}
	// stdlog no-buffer
	for i := 1; i <= 1000; i *= 10 {
		b.Run(fmt.Sprintf("Stdlog-NoBuffer-%d",initN * i), func(b *testing.B) {
			b.ReportAllocs()
			for j := 0; j < b.N; j++ {
				var wg sync.WaitGroup
				wg.Add(initN * i)
				for k := 0; k < initN * i; k++ {
					go func() {
						stdLogger.Println("hello world!")
						wg.Done()
					}()
				}
				wg.Wait()
			}
		})
	}
}
