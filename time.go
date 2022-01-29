package bilog

import (
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	_READER int32 = 0x2
	_WRITER int32 = 0x4
)

type TimeFactory struct {
	buf unsafe.Pointer
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
	timeTmp := time.Now()
	year, month, day := timeTmp.Date()
	hour, minute, second := timeTmp.Hour(), timeTmp.Minute(), timeTmp.Second()
	tmp := fastConvertAllToSlice(year, int(month), day, hour, minute, second)
	atomic.StorePointer(&t.buf, unsafe.Pointer(&tmp))
}

func (t *TimeFactory) Start() {
	t.appendBuf()
	go func() {
		for {
			time.Sleep(time.Millisecond * 10)
			//time.Sleep(time.Second)
			t.appendBuf()
		}
	}()
}

func (t *TimeFactory) Get() []byte {
	//return *(*[32]byte)(unsafe.Pointer(atomic.LoadUintptr((*uintptr)(unsafe.Pointer(&t.buf)))))
	return *(*[]byte)(atomic.LoadPointer(&t.buf))
}

// TODO 废弃
//func (t *TimeFactory) GetRaw() time.Time {
//	for !(atomic.LoadInt32(&t.er) == _READER) {
//		time.Sleep(time.Nanosecond * 2)
//	}
//	return t.raw
//}
