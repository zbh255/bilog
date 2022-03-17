package main

// 测试和验证timeFactory的data-race行为

import (
	"fmt"
	"github.com/zbh255/bilog"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	factory := bilog.NewTimeFactory()
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
		// 收集一万次
		if boolSet[true]+boolSet[false] == 1024 {
			break
		}
	}
	fmt.Println(boolSet)
}
