package concurrentcount

import (
	"sync"
	"testing"
)

const (
	// Number of goroutines using the counter.
	concurrency = 8

	// How many more writes to the counter than reads. In typical telemetry
	// systems, you might have just ~1 read per minute, but possibly
	// millions of reads in the same time.
	writeToReadRatio = 1000

	// Delta used in runAdd.
	delta = 0.5
)

func runInc(b *testing.B, c Counter) {
	b.StopTimer()

	var start, end sync.WaitGroup
	start.Add(1)
	end.Add(concurrency)

	n := b.N / concurrency

	for i := 0; i < concurrency; i++ {
		go func() {
			start.Wait()
			for i := 0; i < n; i++ {
				if i%writeToReadRatio == 0 {
					c.Get()
				} else {
					c.Inc()
				}
			}
			end.Done()
		}()
	}

	b.StartTimer()
	start.Done()
	end.Wait()
}

func runAdd(b *testing.B, c Counter) {
	b.StopTimer()

	var start, end sync.WaitGroup
	start.Add(1)
	end.Add(concurrency)

	n := b.N / concurrency

	for i := 0; i < concurrency; i++ {
		go func() {
			start.Wait()
			for i := 0; i < n; i++ {
				if i%writeToReadRatio == 0 {
					c.Get()
				} else {
					c.Add(delta)
				}
			}
			end.Done()
		}()
	}

	b.StartTimer()
	start.Done()
	end.Wait()
}

func BenchmarkMutexCounterInc(b *testing.B) {
	runInc(b, &MutexCounter{})
}

func BenchmarkMutexCounterAdd(b *testing.B) {
	runAdd(b, &MutexCounter{})
}

func BenchmarkRWMutexCounterInc(b *testing.B) {
	runInc(b, &RWMutexCounter{})
}

func BenchmarkRWMutexCounterAdd(b *testing.B) {
	runAdd(b, &RWMutexCounter{})
}

func BenchmarkAtomicIntCounterInc(b *testing.B) {
	runInc(b, new(AtomicIntCounter))
}

func BenchmarkAtomicIntCounterAdd(b *testing.B) {
	runAdd(b, new(AtomicIntCounter))
}

func BenchmarkAtomicSpinningCounterInc(b *testing.B) {
	runInc(b, new(AtomicSpinningCounter))
}

func BenchmarkAtomicSpinningCounterAdd(b *testing.B) {
	runAdd(b, new(AtomicSpinningCounter))
}

func BenchmarkChannelCounterInc(b *testing.B) {
	runInc(b, NewChannelCounter())
}

func BenchmarkChannelCounterAdd(b *testing.B) {
	runAdd(b, NewChannelCounter())
}
