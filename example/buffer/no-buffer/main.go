package main

import (
	"github.com/zbh255/bilog"
	"os"
)

func main() {
	logger := bilog.NewLogger(os.Stdout,bilog.PANIC,bilog.WithDefault(),
		bilog.WithLowBuffer(0),bilog.WithTopBuffer(0))
	logger.Trace("hello world!")
	logger.Info("hello world!")
}
