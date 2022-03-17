package main

import (
	"github.com/zbh255/bilog"
	"os"
)

func main() {
	logger := bilog.NewLogger(os.Stdout, bilog.PANIC)
	logger.Debug("hello world")
	logger.Trace("hello world!")
	logger.Flush()
}
