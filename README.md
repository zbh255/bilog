# bilog
bilog被设计为可以在需要性能，简单日志的场景下可以替换std log，它关注性能与轻量的实现。

### Benchmark

Cpu: i7-8705G 4C/8T 的测试结果，测试用例在`log_test.go`

```shell
goos: darwin
goarch: amd64
pkg: github.com/zbh255/bilog
cpu: Intel(R) Core(TM) i7-8705G CPU @ 3.10GHz
BenchmarkLogger
BenchmarkLogger/BiLog
BenchmarkLogger/BiLog-8         	11091630	        94.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkLogger/StdLog
BenchmarkLogger/StdLog-8        	 3878521	       319.3 ns/op	      16 B/op	       1 allocs/op
PASS
```

