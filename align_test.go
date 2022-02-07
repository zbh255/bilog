package bilog

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

var simpleLoggerFmt = `
name=mu 		ptr=%p type=%18T offset=%3d size=%3d align=%d
name=factory 	ptr=%p type=%18T offset=%3d size=%3d align=%d
name=level 		ptr=%p type=%18T offset=%3d size=%3d align=%d
name=levelCache ptr=%p type=%18T offset=%3d size=%3d align=%d
name=write 		ptr=%p type=%18T offset=%3d size=%3d align=%d
name=timeBuf 	ptr=%p type=%18T offset=%3d size=%3d align=%d
name=topBuf 	ptr=%p type=%18T offset=%3d size=%3d align=%d
name=lowBuf 	ptr=%p type=%18T offset=%3d size=%3d align=%d
`

func TestLoggerSize(t *testing.T) {
	log := NewLogger(os.Stdout, PANIC)
	fmt.Printf(simpleLoggerFmt, &log.mu, log.mu, unsafe.Offsetof(log.mu), unsafe.Sizeof(&log.mu), unsafe.Alignof(log.mu),
		&log.factory, log.factory, unsafe.Offsetof(log.factory), unsafe.Sizeof(log.factory), unsafe.Alignof(log.factory),
		&log.level, &log.level, unsafe.Offsetof(log.level), unsafe.Sizeof(log.level), unsafe.Alignof(log.level),
		&log.levelCache, &log.levelCache, unsafe.Offsetof(log.levelCache), unsafe.Sizeof(log.levelCache), unsafe.Alignof(log.levelCache),
		&log.write, &log.write, unsafe.Offsetof(log.write), unsafe.Sizeof(log.write), unsafe.Alignof(log.write),
		&log.timeBuf, &log.timeBuf, unsafe.Offsetof(log.timeBuf), unsafe.Sizeof(log.timeBuf), unsafe.Alignof(log.timeBuf),
		&log.topBuf, &log.topBuf, unsafe.Offsetof(log.topBuf), unsafe.Sizeof(log.topBuf), unsafe.Alignof(log.topBuf),
		&log.lowBuf, &log.lowBuf, unsafe.Offsetof(log.lowBuf), unsafe.Sizeof(log.lowBuf), unsafe.Alignof(log.lowBuf))
}
