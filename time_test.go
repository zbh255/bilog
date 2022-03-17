package bilog

import (
	ass "github.com/stretchr/testify/assert"
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
	b.Run("TimeFactoryAppendBuf", func(b *testing.B) {
		b.ReportAllocs()
		factory := NewTimeFactory()
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
		timeBuf := make([]byte, 32)
		oldTimeStamp := time.Now().UnixNano()
		for i := 0; i < b.N; i++ {
			timeStamp := factory.TimeStamp()
			if !(timeStamp-oldTimeStamp > int64(time.Millisecond*10)) {
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
		timeBuf := make([]byte, 32)
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
	return make([]byte, n)
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
		buf[k] = string(tmp)[:len(tmp)-1]
	}
	// 测试时间的误差
	top, err := time.Parse(parse, buf[0])
	if err != nil {
		t.Error(err)
		return
	}
	var offSet int64
	for _, v := range buf {
		tmp, err := time.Parse(parse, v)
		if err != nil {
			t.Error(err)
		}
		offSet = tmp.Unix() - top.Unix()
	}

	if offSet > 10 {
		t.Error("time offSet max")
	}
}

// 比对生成的秒级时间
func TestFactoryCreate(t *testing.T) {
	factory := NewTimeFactory()
	factory.appendBuf()
	// 比对时间戳
	now := time.Now()
	assert := ass.New(t)
	assert.Equal(now.Unix(),factory.TimeStamp()/int64(time.Second))
	// 比对序列化时间
	assert.Equal(factory.Get(),fastConvertAllToSlice(now.Year(), int(now.Month()),
		now.Day(),now.Hour(),
		now.Minute(),now.Second(),
	))
}

// panic日志的测试
func TestFactoryConcurrentCreate(t *testing.T) {
	assert := ass.New(t)
	factory := NewTimeFactory()
	factory.Start()
	timeBuf := string(factory.Get())
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		newTimeBuf := string(factory.Get())
		assert.NotEqual(timeBuf,newTimeBuf)
	}
}

func TestFactoryTimeStamp(t *testing.T) {
	factory := NewTimeFactory()
	factory.Start()
	// 间隔0.01ms收集factory的bool信息
	boolSet := make(map[bool]int, 10)
	var oldTimeStamp = time.Now().UnixNano()
	for {
		time.Sleep(time.Millisecond * 5)
		timeStamp := factory.TimeStamp()
		if timeStamp-oldTimeStamp > int64(time.Millisecond*10) {
			boolSet[true]++
		} else {
			boolSet[false]++
		}
		oldTimeStamp = timeStamp
		// 收集1024次
		if boolSet[true]+boolSet[false] == 1024 {
			break
		}
	}
	assert := ass.New(t)
	assert.NotEqual(boolSet[false],0)
}