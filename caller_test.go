package bilog

import (
	"testing"
)

func TestCallerFeatures(t *testing.T) {
	for i := 0; i < 5;i++ {
		_,_ = Caller(3)
		_,_ = CallerOfCache(3)
		_,_ = CallerOfConcurrentCache(3)
	}
}

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


