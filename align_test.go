package bilog

import (
	"os"
	"testing"
	"unsafe"
)



func TestLoggerSize(t *testing.T) {
	log := NewLogger(os.Stdout, PANIC)
	t.Log(unsafe.Sizeof(*log))
}
