package bilog

import (
	"regexp"
	"testing"
)

func TestCallerFeatures(t *testing.T) {
	for i := 0; i < 5; i++ {
		_, _ = Caller(3)
		_, _ = CallerOfCache(3)
		_, _ = CallerOfConcurrentCache(3)
	}
}

func TestCutString(t *testing.T) {
	compile, err := regexp.Compile("bilog[/\\\\]{1,}logger.go")
	if err != nil {
		t.Error(err)
		return
	}
	if !compile.MatchString(cutFileName("/Users/harder/github.com/nyan233/bilog/logger.go")) {
		t.Error("format failed")
	}
	if !compile.MatchString(cutFileName("\\Users\\github.com\\nyan233\\bilog\\logger.go")) {
		t.Error("format failed")
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
			_, _ = CallerOfConcurrentCache(3)
		}
	})
}
