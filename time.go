package bilog

import (
	"sync"
	"time"
)

type TimeFactory struct {
	rwMu sync.RWMutex
	buf  *date
	raw  time.Time
}

type date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

func (t *TimeFactory) Start() {
	if t.buf == nil {
		t.raw = time.Now()
		t.buf = &date{
			Year:   t.raw.Year(),
			Month:  int(t.raw.Month()),
			Day:    t.raw.Day(),
			Hour:   t.raw.Hour(),
			Minute: t.raw.Minute(),
			Second: t.raw.Second(),
		}
	}
	go func() {
		for {
			time.Sleep(time.Second)
			t.rwMu.Lock()
			t.raw = time.Now()
			t.buf.Year = t.raw.Year()
			t.buf.Month = int(t.raw.Month())
			t.buf.Day = t.raw.Day()
			t.buf.Hour = t.raw.Hour()
			t.buf.Minute = t.raw.Minute()
			t.buf.Second = t.raw.Second()
			t.rwMu.Unlock()
		}
	}()
}

func (t *TimeFactory) Get() *date {
	t.rwMu.RLock()
	defer t.rwMu.RUnlock()
	return t.buf
}

func (t *TimeFactory) GetRaw() time.Time {
	t.rwMu.RLock()
	defer t.rwMu.RUnlock()
	return t.raw
}
