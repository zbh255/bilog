package bilog

import (
	"sync/atomic"
	"time"
)

const (
	_READER int32 = 0x2
	_WRITER int32 = 0x4
)

type TimeFactory struct {
	eff int
	buf atomic.Value
	raw time.Time
}

func NewTimeFactory() *TimeFactory {
	return &TimeFactory{}
}

//func (t *TimeFactory) appendBuf() {
//	// reset
//	h := (*reflect.SliceHeader)(unsafe.Pointer(&t.buf))
//	h.Len = 0
//	t.buf = *(*[]byte)(unsafe.Pointer(h))
//	// append
//	t.buf = append(t.buf, fastConvertYear(t.raw.Year())...)
//	t.buf = append(t.buf, fastConvertMonth(int(t.raw.Month()))...)
//	t.buf = append(t.buf, fastConvertDay(t.raw.Day())...)
//	t.buf = append(t.buf, fastConvertHour(t.raw.Hour())...)
//	t.buf = append(t.buf, fastConvertMinute(t.raw.Minute())...)
//	t.buf = append(t.buf, fastConvertSecond(t.raw.Second())...)
//}

func (t *TimeFactory) appendBuf() {
	t.raw = time.Now()
	year, month, day := t.raw.Date()
	hour, minute, second := t.raw.Hour(), t.raw.Minute(), t.raw.Second()
	tmp := fastConvertAllToSlice(year, int(month), day, hour, minute, second)
	t.buf.Store(tmp)
}

func (t *TimeFactory) Start() {
	t.appendBuf()
	go func() {
		for {
			time.Sleep(time.Millisecond * 10)
			t.appendBuf()
		}
	}()
}

func (t *TimeFactory) Get() []byte {
	//return *(*[32]byte)(unsafe.Pointer(atomic.LoadUintptr((*uintptr)(unsafe.Pointer(&t.buf)))))
	return t.buf.Load().([]byte)
}

// TODO 废弃
//func (t *TimeFactory) GetRaw() time.Time {
//	for !(atomic.LoadInt32(&t.er) == _READER) {
//		time.Sleep(time.Nanosecond * 2)
//	}
//	return t.raw
//}
