package bilog

import (
	"testing"
	"time"
)

func BenchmarkTimeFactory(b *testing.B) {
	b.Run("NoCache", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = time.Now()
		}
	})
	b.Run("TimeFactoryRaw", func(b *testing.B) {
		b.ReportAllocs()
		factory := &TimeFactory{}
		factory.Start()
		for i := 0; i < b.N; i++ {
			_ = factory.GetRaw()
		}
	})
	b.Run("TimeFactoryNoRaw", func(b *testing.B) {
		b.ReportAllocs()
		factory := &TimeFactory{}
		factory.Start()
		for i := 0; i < b.N; i++ {
			_ = factory.Get()
		}
	})
}

// 测试生成的时间序列是否连续
func TestTimeFactory(t *testing.T) {
	factory := &TimeFactory{}
	factory.Start()
	buf := make([]time.Time, 100)
	for k := range buf {
		time.Sleep(time.Millisecond * 10)
		buf[k] = factory.GetRaw()
	}
	// 测试时间的误差
	top := buf[0].Unix()
	var offSet int64
	for _, v := range buf {
		offSet = v.Unix() - top
	}
	if offSet > 10 {
		t.Error("time offSet max")
	}
}
