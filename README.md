# Bilog [![Go Report Card](https://goreportcard.com/badge/github.com/zbh255/bilog)](https://goreportcard.com/report/github.com/zbh255/bilog) ![GitHub](https://img.shields.io/github/license/zbh255/bilog) ![GitHub](https://github.com/zbh255/bilog/actions/workflows/go.yml/badge.svg) [![Go Doc](https://pkg.go.dev/badge/github.com/zbh255/bilog?utm_source=godoc)](https://pkg.go.dev/github.com/zbh255/bilog)

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
pkg: github.com/zbh255/bilog/benchmark
cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
BenchmarkNoConcurrent
BenchmarkNoConcurrent/Bilog-Buffer-1000
BenchmarkNoConcurrent/Bilog-Buffer-1000-8         	   28856	     42290 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-Buffer-10000
BenchmarkNoConcurrent/Bilog-Buffer-10000-8        	    3130	    384731 ns/op	       1 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-Buffer-100000
BenchmarkNoConcurrent/Bilog-Buffer-100000-8       	     307	   3811345 ns/op	      18 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-1000
BenchmarkNoConcurrent/Bilog-NoBuffer-1000-8       	   31294	     36844 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-10000
BenchmarkNoConcurrent/Bilog-NoBuffer-10000-8      	    3140	    370229 ns/op	       1 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-100000
BenchmarkNoConcurrent/Bilog-NoBuffer-100000-8     	     321	   3677432 ns/op	      17 B/op	       0 allocs/op
BenchmarkNoConcurrent/Bilog-NoBuffer-1000000
BenchmarkNoConcurrent/Bilog-NoBuffer-1000000-8    	      32	  37684024 ns/op	     183 B/op	       6 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000-8      	    4122	    285163 ns/op	   16006 B/op	    1000 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-10000
BenchmarkNoConcurrent/Stdlog-NoBuffer-10000-8     	     420	   2877181 ns/op	  160072 B/op	   10000 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-100000
BenchmarkNoConcurrent/Stdlog-NoBuffer-100000-8    	      39	  28613590 ns/op	 1600722 B/op	  100007 allocs/op
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000000
BenchmarkNoConcurrent/Stdlog-NoBuffer-1000000-8   	       4	 286630937 ns/op	16006828 B/op	 1000072 allocs/op
BenchmarkConcurrent
BenchmarkConcurrent/Bilog-Buffer-1000
BenchmarkConcurrent/Bilog-Buffer-1000-8           	    3691	    329624 ns/op	   24218 B/op	    1002 allocs/op
BenchmarkConcurrent/Bilog-Buffer-10000
BenchmarkConcurrent/Bilog-Buffer-10000-8          	     427	   2901827 ns/op	  252257 B/op	   10117 allocs/op
BenchmarkConcurrent/Bilog-Buffer-100000
BenchmarkConcurrent/Bilog-Buffer-100000-8         	      38	  31092050 ns/op	 2781661 B/op	  102667 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-1000
BenchmarkConcurrent/Bilog-NoBuffer-1000-8         	    3998	    322535 ns/op	   24090 B/op	    1001 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-10000
BenchmarkConcurrent/Bilog-NoBuffer-10000-8        	     414	   2804408 ns/op	  243948 B/op	   10042 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-100000
BenchmarkConcurrent/Bilog-NoBuffer-100000-8       	      43	  29338783 ns/op	 2539998 B/op	  101435 allocs/op
BenchmarkConcurrent/Bilog-NoBuffer-1000000
BenchmarkConcurrent/Bilog-NoBuffer-1000000-8      	       4	 290937219 ns/op	25162768 B/op	 1012153 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-1000
BenchmarkConcurrent/Stdlog-NoBuffer-1000-8        	     428	   2874261 ns/op	   40569 B/op	    2006 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-10000
BenchmarkConcurrent/Stdlog-NoBuffer-10000-8       	      40	  28370353 ns/op	  434488 B/op	   20361 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-100000
BenchmarkConcurrent/Stdlog-NoBuffer-100000-8      	       4	 273002783 ns/op	 6138152 B/op	  222075 allocs/op
BenchmarkConcurrent/Stdlog-NoBuffer-1000000
BenchmarkConcurrent/Stdlog-NoBuffer-1000000-8     	       1	3035376571 ns/op	418629608 B/op	 3412337 allocs/op
PASS
```

## Lisence

The Bilog Use Mit licensed. More is See [Lisence](https://github.com/zbh255/bilog/blob/main/LICENSE)

