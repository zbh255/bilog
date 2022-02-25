# Bilog [![Go Report Card](https://goreportcard.com/badge/github.com/zbh255/bilog)](https://goreportcard.com/report/github.com/zbh255/bilog) ![GitHub](https://img.shields.io/github/license/zbh255/bilog) ![GitHub](https://github.com/zbh255/bilog/actions/workflows/go.yml/badge.svg)

bilog被设计为可以在需要性能，简单日志的场景下可以替换std log，它关注性能与轻量的实现。

## Install

```shell
go get github.com/zbh255/bilog
```

## Usage

```go
func main() {
	logger := bilog.NewLogger(os.Stdout, bilog.PANIC)
	logger.Debug("hello world")
	logger.Trace("hello world!")
	logger.Flush()
}
```

`OutPut`

```shell
[DEBUG] 2022-01-24 12:53:29 hello world
[TRACE] 2022-01-24 12:53:29 hello world!
```

### Benchmark

Cpu: i7-8705G 4C/8T 的测试结果，测试用例在`log_test.go`

```shell
goos: darwin
goarch: amd64
pkg: github.com/zbh255/bilog
cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
BenchmarkLogger
BenchmarkLogger/BiLog
BenchmarkLogger/BiLog-8         	24042242	        43.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkLogger/BiLogDouble
BenchmarkLogger/BiLogDouble-8   	14993350	        80.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkLogger/BilogCallerAndTime
BenchmarkLogger/BilogCallerAndTime-8         	 1952274	       608.8 ns/op	     224 B/op	       3 allocs/op
BenchmarkLogger/StdLogCallerAndTime
BenchmarkLogger/StdLogCallerAndTime-8        	 1482902	       804.7 ns/op	     232 B/op	       3 allocs/op
BenchmarkLogger/StdLog
BenchmarkLogger/StdLog-8                     	 4112342	       288.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkLogger/StdLogDouble
BenchmarkLogger/StdLogDouble-8               	 1974378	       604.7 ns/op	      32 B/op	       2 allocs/op
PASS
```

## Lisence

The Bilog Use Mit licensed. More is See [Lisence](https://github.com/zbh255/bilog/blob/main/LICENSE)

