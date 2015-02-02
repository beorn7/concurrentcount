# concurrentcount

Experiments to benchmark implementations of a concurrent counter in Go.

Such a counter is useful to export telemetry. See
[Prometheus](http://prometheus.io) as an example for a powerful
monitoring system.

You can configure concurrency and other parameters in the
`benchmark_test.go` file directly.

To run the benchmarks with different values for `GOMAXPROCS`:

`go test -bench . -benchmem -cpu 1,2,4,8,16`


Example output:

```
BenchmarkMutexCounterInc        20000000               114 ns/op
BenchmarkMutexCounterInc-2       5000000               321 ns/op
BenchmarkMutexCounterInc-4       5000000               376 ns/op
BenchmarkMutexCounterInc-8       2000000               547 ns/op
BenchmarkMutexCounterInc-16      2000000               603 ns/op
BenchmarkMutexCounterAdd        20000000               114 ns/op
BenchmarkMutexCounterAdd-2       5000000               328 ns/op
BenchmarkMutexCounterAdd-4       5000000               376 ns/op
BenchmarkMutexCounterAdd-8       5000000               349 ns/op
BenchmarkMutexCounterAdd-16      5000000               427 ns/op
BenchmarkRWMutexCounterInc      10000000               130 ns/op
BenchmarkRWMutexCounterInc-2     3000000               453 ns/op
BenchmarkRWMutexCounterInc-4     3000000               493 ns/op
BenchmarkRWMutexCounterInc-8     3000000               664 ns/op
BenchmarkRWMutexCounterInc-16    2000000               709 ns/op
BenchmarkRWMutexCounterAdd      10000000               131 ns/op
BenchmarkRWMutexCounterAdd-2     3000000               426 ns/op
BenchmarkRWMutexCounterAdd-4     3000000               524 ns/op
BenchmarkRWMutexCounterAdd-8     2000000               690 ns/op
BenchmarkRWMutexCounterAdd-16    2000000               648 ns/op
BenchmarkAtomicIntCounterInc    100000000               10.5 ns/op
BenchmarkAtomicIntCounterInc-2  30000000                37.7 ns/op
BenchmarkAtomicIntCounterInc-4  50000000                39.5 ns/op
BenchmarkAtomicIntCounterInc-8  50000000                33.6 ns/op
BenchmarkAtomicIntCounterInc-16 50000000                34.5 ns/op
BenchmarkAtomicIntCounterAdd    100000000               10.6 ns/op
BenchmarkAtomicIntCounterAdd-2  50000000                37.2 ns/op
BenchmarkAtomicIntCounterAdd-4  50000000                36.4 ns/op
BenchmarkAtomicIntCounterAdd-8  50000000                35.0 ns/op
BenchmarkAtomicIntCounterAdd-16 50000000                35.3 ns/op
BenchmarkAtomicSpinningCounterInc       100000000               16.5 ns/op
BenchmarkAtomicSpinningCounterInc-2     20000000                68.8 ns/op
BenchmarkAtomicSpinningCounterInc-4     20000000               118 ns/op
BenchmarkAtomicSpinningCounterInc-8     10000000               180 ns/op
BenchmarkAtomicSpinningCounterInc-16    10000000               175 ns/op
BenchmarkAtomicSpinningCounterAdd       100000000               16.6 ns/op
BenchmarkAtomicSpinningCounterAdd-2     20000000                68.7 ns/op
BenchmarkAtomicSpinningCounterAdd-4     20000000               119 ns/op
BenchmarkAtomicSpinningCounterAdd-8     10000000               179 ns/op
BenchmarkAtomicSpinningCounterAdd-16    10000000               177 ns/op
BenchmarkChannelCounterInc      20000000                73.3 ns/op
BenchmarkChannelCounterInc-2    10000000               164 ns/op
BenchmarkChannelCounterInc-4     5000000               333 ns/op
BenchmarkChannelCounterInc-8     3000000               463 ns/op
BenchmarkChannelCounterInc-16    3000000               472 ns/op
BenchmarkChannelCounterAdd      20000000                73.3 ns/op
BenchmarkChannelCounterAdd-2    10000000               152 ns/op
BenchmarkChannelCounterAdd-4     5000000               290 ns/op
BenchmarkChannelCounterAdd-8     3000000               473 ns/op
BenchmarkChannelCounterAdd-16    3000000               476 ns/op
```
