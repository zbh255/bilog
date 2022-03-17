package bilog

import (
	"sync/atomic"
	"time"
)

type TimeFactoryZero struct {
	buf atomic.Value
}

type Array struct {
	buf [32]byte
	eff int
}

func NewTimeFactoryZero() *TimeFactoryZero {
	return &TimeFactoryZero{}
}

func (t *TimeFactoryZero) appendBuf() {
	timeTmp := time.Now()
	year, month, day := timeTmp.Date()
	hour, minute, second := timeTmp.Hour(), timeTmp.Minute(), timeTmp.Second()
	tmp, eff := fastConvertAllToArray(year, int(month), day, hour, minute, second)
	t.buf.Store(Array{
		buf: tmp,
		eff: eff,
	})
}

func (t *TimeFactoryZero) Start() {
	t.appendBuf()
	go func() {
		for {
			time.Sleep(time.Millisecond * 10)
			t.appendBuf()
		}
	}()
}

func (t *TimeFactoryZero) Get() ([32]byte, int) {
	//return *(*[32]byte)(unsafe.Pointer(atomic.LoadUintptr((*uintptr)(unsafe.Pointer(&t.buf)))))
	tmp := t.buf.Load().(Array)
	return tmp.buf, tmp.eff
}
