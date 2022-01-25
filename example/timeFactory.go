package main

import "github.com/zbh255/bilog"

func main() {
	factory := bilog.NewTimeFactory()
	factory.Start()
	factory.Get()
}
