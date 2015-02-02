# concurrentcount

Experiments to benchmark implementations of a concurrent counter in Go.

Such a counter is useful to export telemetry. See
[Prometheus](http://prometheus.io) as an example for a powerful
monitoring system.

You can configure concurrency and other parameters in the
`benchmark_test.go` file directly.

To run the benchmarks with different values for `GOMAXPROCS`:

`go test -bench . -benchmem -cpu 1,2,4,8,16`
