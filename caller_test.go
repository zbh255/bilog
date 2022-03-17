package bilog

import (
	"testing"
)

func BenchmarkCaller(b *testing.B) {
	b.Run("Default", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = Caller(3)
		}
	})
	b.Run("CallerCached", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = CallerOfCache(3)
		}
	})
	b.Run("CallerConcurrentCache", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_,_ = CallerOfConcurrentCache(3)
		}
	})
}


