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
	b.Run("TimeFactoryNoRaw", func(b *testing.B) {
		b.ReportAllocs()
		factory := &TimeFactory{}
		factory.Start()
		for i := 0; i < b.N; i++ {
			_ = factory.Get()
		}
	})
}

// 禁止编译器优化: -gcflags "-N -l"
/*
goos: darwin
goarch: amd64
pkg: github.com/zbh255/bilog
cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
BenchmarkMemoryCreate
BenchmarkMemoryCreate/Heap
BenchmarkMemoryCreate/Heap-8         	52724396	        23.65 ns/op	      32 B/op	       1 allocs/op
BenchmarkMemoryCreate/Stack
BenchmarkMemoryCreate/Stack-8        	482789616	         2.459 ns/op	       0 B/op	       0 allocs/op
PASS
*/
func BenchmarkMemoryCreate(b *testing.B) {
	b.Run("Heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = heapCreate(32)
		}
	})
	b.Run("Stack", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = stackCreate()
		}
	})
}

func heapCreate(n int) []byte {
	return make([]byte,n)
}

func stackCreate() [32]byte {
	var tmp [32]byte
	return tmp
}

// 测试生成的时间序列是否连续
func TestTimeFactory(t *testing.T) {
	factory := &TimeFactory{}
	factory.Start()
	buf := make([][]byte, 100)
	for k := range buf {
		time.Sleep(time.Millisecond * 10)
		tmp := factory.Get()
		buf[k] = tmp
	}
	// 测试时间的误差
	//top := buf[0].Unix()
	//var offSet int64
	//for _, v := range buf {
	//	offSet = v.Unix() - top
	//}
	//if offSet > 10 {
	//	t.Error("time offSet max")
	//}
}

func TestFactoryCreate(t *testing.T) {
	factory := NewTimeFactory()
	factory.Start()
	t.Log(factory.Get())
}