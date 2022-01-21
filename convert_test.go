package bilog

import (
	"strconv"
	"testing"
	"time"
)

func BenchmarkConvert(b *testing.B) {
	b.Run("Strconv", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			yy, mm, dd := time.Now().Date()
			_ = strconv.Itoa(yy)
			_ = strconv.Itoa(int(mm))
			_ = strconv.Itoa(dd)
		}
	})
	b.Run("FastConvert", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			yy, mm, dd := time.Now().Date()
			_ = fastConvertYear(yy)
			_ = fastConvertMonth(int(mm))
			_ = fastConvertDay(dd)
		}
	})
	b.Run("FastConvertAndFactory", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		factory := &TimeFactory{}
		factory.Start()
		for i := 0; i < b.N; i++ {
			date := factory.GetRaw()
			_ = fastConvertYear(date.Year())
			_ = fastConvertMonth(int(date.Month()))
			_ = fastConvertDay(date.Day())
		}
	})
}
