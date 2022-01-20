package bilog

import (
	"sync"
	"time"
)

type TimeFactory struct {
	rwMu sync.RWMutex
	buf  *time.Time
}

func (t *TimeFactory) Start() {
	if t.buf == nil {
		tmp := time.Now()
		t.buf = &tmp
	}
	go func() {
		for {
			time.Sleep(time.Millisecond)
			t.rwMu.Lock()
			*t.buf = time.Now()
			t.rwMu.Unlock()
		}
	}()
}

func (t *TimeFactory) Get() time.Time {
	t.rwMu.RLock()
	defer t.rwMu.RUnlock()
	return *t.buf
}
