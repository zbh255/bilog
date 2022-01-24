package main

import (
	"github.com/zbh255/bilog"
	"os"
	"time"
)

func main() {
	logger := bilog.NewLogger(os.Stdout, bilog.ERROR)
	logger.Debug("hello world")
	time.Sleep(time.Second)
	logger.Trace("hello world!")
	logger.Flush()
}
