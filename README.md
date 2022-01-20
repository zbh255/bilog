# bilog
bilog被设计为可以在需要性能，简单日志的场景下可以替换std log，它关注性能与轻量的实现。

### Benchmark

Cpu: i7-8705G 4C/8T 的测试结果，测试用例在`log_test.go`

```shell
BenchmarkLogger
BenchmarkLogger/BiLog
BenchmarkLogger/BiLog-8         	17665779	        63.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkLogger/BiLogSwitchPrefix
BenchmarkLogger/BiLogSwitchPrefix-8         	13847743	        73.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkLogger/StdLog
BenchmarkLogger/StdLog-8                    	 3909732	       303.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkLogger/StdLogSwitchPrefix
BenchmarkLogger/StdLogSwitchPrefix-8        	 3709778	       313.6 ns/op	      16 B/op	       1 allocs/op
PASS

```

