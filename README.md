# Bilog [![Go Report Card](https://goreportcard.com/badge/github.com/zbh255/bilog)](https://goreportcard.com/report/github.com/zbh255/bilog) ![GitHub](https://img.shields.io/github/license/zbh255/bilog) ![GitHub](https://github.com/zbh255/bilog/actions/workflows/go.yml/badge.svg) [![codecov](https://codecov.io/gh/zbh255/bilog/branch/main/graph/badge.svg?token=DKaQWjgsF8)](https://codecov.io/gh/zbh255/bilog) [![Go Doc](https://pkg.go.dev/badge/github.com/zbh255/bilog?utm_source=godoc)](https://pkg.go.dev/github.com/zbh255/bilog)

bilog被设计为可以在需要性能，简单日志的场景下可以替换std log，它关注性能与轻量的实现。

## Install

```shell
go get github.com/zbh255/bilog
```

## Quick-Start

> 下面的所有示例代码都可在本`repo`的`example`文件夹中找到

### Print-Time

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

### Print-Caller

```go
func main() {
	logger := bilog.NewLogger(os.Stdout,bilog.PANIC,bilog.WithDefault(),bilog.WithCaller(0))
	logger.Trace("hello world!")
	logger.Debug("hello world!")
	logger.Flush()
}
```

`OutPut`

```shell
[TRACE] 2022-05-07 23:57:27 /Users/harder/Desktop/Git-Repo/github.com/zbh255/bilog/example/basic/caller/main.go:10 hello world!
[DEBUG] 2022-05-07 23:57:27 /Users/harder/Desktop/Git-Repo/github.com/zbh255/bilog/example/basic/caller/main.go:11 hello world!
```

### No-Buffer

`bilog`默认使用双重缓冲区来缓冲需要打印的`bytes`，您可以像上面的例子那样使用`Flush`强制刷新`Buffer`，也可以禁用缓冲，如下所示。

```go
func main() {
	logger := bilog.NewLogger(os.Stdout,bilog.PANIC,bilog.WithDefault(),
		bilog.WithLowBuffer(0),bilog.WithTopBuffer(0))
	logger.Trace("hello world!")
	logger.Info("hello world!")
}
```

`OutPut`

```shell
[TRACE] 2022-05-08 00:35:34  hello world!
[INFO] 2022-05-08 00:35:34  hello world!
```

> 事实上，您可以禁用缓冲也可以自行调节缓冲区的大小，使用跟如上示例一样的`Api`

## Benchmark

Cpu: i7-8705G 4C/8T 的测试结果，测试用例在`log_test.go`

```shell
goos: darwin
goarch: amd64
pkg: github.com/zbh255/bilog/benchmark
cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
BenchmarkNoConcurrent
BenchmarkNoConcurrent/Bilog-Buffer-1000
BenchmarkNoConcurrent/Bilog-Buffer-1000-8         	   30687	     38321 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-Buffer-10000
BenchmarkNoConcurrent/Bilog-Buffer-10000-8        	    3160	    371820 ns/op	       1 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-Buffer-100000
BenchmarkNoConcurrent/Bilog-Buffer-100000-8       	     328	   3586913 ns/op	      17 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-Buffer-1000000
BenchmarkNoConcurrent/Bilog-Buffer-1000000-8      	      32	  35836612 ns/op	     173 B/op	       6 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-1000
BenchmarkNoConcurrent/Bilog-NoBuffer-1000-8       	   37311	     31553 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-10000
BenchmarkNoConcurrent/Bilog-NoBuffer-10000-8      	    3680	    317317 ns/op	       1 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-100000
BenchmarkNoConcurrent/Bilog-NoBuffer-100000-8     	     376	   3170883 ns/op	      15 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-1000000
BenchmarkNoConcurrent/Bilog-NoBuffer-1000000-8    	      34	  31894623 ns/op	     156 B/op	       5 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000-8      	    3900	    287960 ns/op	   16007 B/op	    1000 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-10000
BenchmarkNoConcurrent/Stdlog-NoBuffer-10000-8     	     382	   2857697 ns/op	  160069 B/op	   10000 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-100000
BenchmarkNoConcurrent/Stdlog-NoBuffer-100000-8    	      40	  28701547 ns/op	 1600674 B/op	  100007 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000000
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000000-8   	       4	 286931206 ns/op	16006676 B/op	 1000070 allocs/op
BenchmarkConcurrent
BenchmarkConcurrent/Bilog-Buffer-1000
BenchmarkConcurrent/Bilog-Buffer-1000-8           	    3792	    323843 ns/op	   24203 B/op	    1002 allocs/op
BenchmarkConcurrent/Bilog-Buffer-10000
BenchmarkConcurrent/Bilog-Buffer-10000-8          	     458	   2735214 ns/op	  249576 B/op	   10094 allocs/op
BenchmarkConcurrent/Bilog-Buffer-100000
BenchmarkConcurrent/Bilog-Buffer-100000-8         	      42	  28385323 ns/op	 2661977 B/op	  101717 allocs/op
BenchmarkConcurrent/Bilog-Buffer-1000000
BenchmarkConcurrent/Bilog-Buffer-1000000-8        	       4	 279774264 ns/op	25004370 B/op	 1008890 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-1000
BenchmarkConcurrent/Bilog-NoBuffer-1000-8         	    4624	    286396 ns/op	   24065 B/op	    1001 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-10000
BenchmarkConcurrent/Bilog-NoBuffer-10000-8        	     494	   2411639 ns/op	  242474 B/op	   10026 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-100000
BenchmarkConcurrent/Bilog-NoBuffer-100000-8       	      54	  23822805 ns/op	 2461564 B/op	  100645 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-1000000
BenchmarkConcurrent/Bilog-NoBuffer-1000000-8      	       5	 233709438 ns/op	24369846 B/op	 1003884 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-1000
BenchmarkConcurrent/Stdlog-NoBuffer-1000-8        	     429	   2856284 ns/op	   40511 B/op	    2006 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-10000
BenchmarkConcurrent/Stdlog-NoBuffer-10000-8       	      38	  27950113 ns/op	  439271 B/op	   20407 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-100000
BenchmarkConcurrent/Stdlog-NoBuffer-100000-8      	       4	 262928956 ns/op	 6094424 B/op	  221852 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-1000000
BenchmarkConcurrent/Stdlog-NoBuffer-1000000-8     	       1	3148452471 ns/op	420471352 B/op	 3423558 allocs/op
PASS
```

## Lisence

The Bilog Use Mit licensed. More is See [Lisence](https://github.com/zbh255/bilog/blob/main/LICENSE)

