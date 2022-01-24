package bilog

import (
	"reflect"
	"sync"
	"time"
	"unsafe"
)

type TimeFactory struct {
	rwMu sync.RWMutex
	buf  []byte
	raw  time.Time
}

func NewTimeFactory() *TimeFactory {
	return &TimeFactory{
		buf: make([]byte, 0, TIME_BUF_SIZE),
	}
}

func (t *TimeFactory) appendBuf() {
	// reset
	h := (*reflect.SliceHeader)(unsafe.Pointer(&t.buf))
	h.Len = 0
	t.buf = *(*[]byte)(unsafe.Pointer(h))
	// append
	t.buf = append(t.buf, fastConvertYear(t.raw.Year())...)
	t.buf = append(t.buf, fastConvertMonth(int(t.raw.Month()))...)
	t.buf = append(t.buf, fastConvertDay(t.raw.Day())...)
	t.buf = append(t.buf, fastConvertHour(t.raw.Hour())...)
	t.buf = append(t.buf, fastConvertMinute(t.raw.Minute())...)
	t.buf = append(t.buf, fastConvertSecond(t.raw.Second())...)
}

func (t *TimeFactory) Start() {
	if t.buf == nil || len(t.buf) == 0 {
		t.raw = time.Now()
		t.appendBuf()
	}
	go func() {
		for {
			time.Sleep(time.Millisecond * 10)
			t.rwMu.Lock()
			t.raw = time.Now()
			t.appendBuf()
			t.rwMu.Unlock()
		}
	}()
}

func (t *TimeFactory) Get() []byte {
	t.rwMu.RLock()
	defer t.rwMu.RUnlock()
	return t.buf
}

func (t *TimeFactory) GetRaw() time.Time {
	t.rwMu.RLock()
	defer t.rwMu.RUnlock()
	return t.raw
}
