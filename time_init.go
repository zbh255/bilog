package bilog

import (
	"fmt"
	"runtime"
	"time"
)

var timeFactory *TimeFactory

func init() {
	timeFactory = NewTimeFactory()
	go func() {
		for {
			time.Sleep(time.Second / 10)
			fmt.Println(runtime.NumGoroutine())
		}
	}()
}
