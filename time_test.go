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
		factory := NewTimeFactory()
		factory.Start()
		for i := 0; i < b.N; i++ {
			_ = factory.Get()
		}
	})
	b.Run("TimeFactoryZeroNoRaw", func(b *testing.B) {
		b.ReportAllocs()
		factory := NewTimeFactoryZero()
		factory.Start()
		for i := 0; i < b.N; i++ {
			_, _ = factory.Get()
		}
	})
	b.Run("TimeFactoryAppendBuf", func(b *testing.B) {
		b.ReportAllocs()
		factory := NewTimeFactory()
		for i := 0; i < b.N; i++ {
			factory.appendBuf()
		}
	})
	b.Run("TimeFactoryZeroAppendBuf", func(b *testing.B) {
		b.ReportAllocs()
		factory := NewTimeFactoryZero()
		for i := 0; i < b.N; i++ {
			factory.appendBuf()
		}
	})
}

func BenchmarkTimeFactoryUpdateOf(b *testing.B) {
	b.Run("UsageUpdateOf", func(b *testing.B) {
		b.ReportAllocs()
		factory := NewTimeFactory()
		factory.Start()
		timeBuf := make([]byte,32)
		oldTimeStamp := time.Now().UnixNano()
		for i := 0; i < b.N; i++ {
			timeStamp := factory.TimeStamp()
			if !(timeStamp - oldTimeStamp > int64(time.Millisecond * 10)) {
				continue
			} else {
				copy(timeBuf[:], factory.Get())
			}
			oldTimeStamp = timeStamp
		}
	})
	b.Run("NoUsageUpdateOf", func(b *testing.B) {
		b.ReportAllocs()
		factory := NewTimeFactory()
		factory.Start()
		timeBuf := make([]byte,32)
		for i := 0; i < b.N; i++ {
			copy(timeBuf[:], factory.Get())
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
	parse := "2006-01-02 15:04:05"
	factory := &TimeFactory{}
	factory.Start()
	buf := make([]string, 100)
	for k := range buf {
		time.Sleep(time.Millisecond * 10)
		tmp := factory.Get()
		buf[k] = string(tmp)[:len(tmp) - 1]
	}
	// 测试时间的误差
	top,err := time.Parse(parse,buf[0])
	if err != nil {
		t.Error(err)
		return
	}
	var offSet int64
	for _, v := range buf {
		tmp,err := time.Parse(parse,v)
		if err != nil {
			t.Error(err)
		}
		offSet = tmp.Unix() - top.Unix()
	}

	if offSet > 10 {
		t.Error("time offSet max")
	}
}

func TestFactoryCreate(t *testing.T) {
	factory := NewTimeFactory()
	factory.Start()
	t.Log(factory.Get())
}


// panic日志的测试
// fatal error: found bad pointer in Go heap (incorrect use of unsafe or cgo?)
func TestFactoryPointer(t *testing.T) {
	factory := NewTimeFactory()
	factory.Start()
	factory.Get()
	time.Sleep(time.Second * 10)
}