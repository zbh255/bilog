package main

import (
	"errors"
	"github.com/zbh255/bilog"
	"os"
)

func main() {
	log := bilog.NewLogger(os.Stdout, bilog.PANIC,
		bilog.WithTimes(), bilog.WithCaller(0), bilog.WithTopBuffer(7), bilog.WithLowBuffer(2))
	log.Info("buffer1")
	log.Info("buffer 2")
	log.Debug("buffer 3")
	log.ErrorFromString("buffer 4")
	log.ErrorFromErr(errors.New("My Error: buffer 5"))
}
