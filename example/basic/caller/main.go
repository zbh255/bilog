package main

import (
	"github.com/zbh255/bilog"
	"os"
)

func main() {
	logger := bilog.NewLogger(os.Stdout,bilog.PANIC,bilog.WithDefault(),bilog.WithCaller())
	logger.Trace("hello world!")
	logger.Debug("hello world!")
	logger.Flush()
}
