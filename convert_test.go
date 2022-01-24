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

// 测试转换的正确性
func TestFastConvert(t *testing.T) {
	createTimeSequence()
}

func createTimeSequence() []string {
	sequence := make([]string, 0, 60*60*24*31*12)
	for year := 1970; year <= 1970; year++ {
		for month := 1; month <= 12; month++ {
			for day := 1; day <= 31; day++ {
				for hour := 0; hour < 24; hour++ {
					for minute := 0; minute < 60; minute++ {
						for second := 0; second < 60; second++ {
							str := fastConvertYear(year) +
								fastConvertMonth(month) +
								fastConvertDay(day) +
								fastConvertHour(hour) +
								fastConvertMinute(minute)
							fastConvertSecond(second)
							sequence = append(sequence, str)
						}
					}
				}
			}
		}
	}
	return sequence
}
