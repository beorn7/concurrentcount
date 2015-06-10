package concurrentcount

import (
	"sync"
	"testing"
)

const (
	// How many more writes to the counter than reads. In typical telemetry
	// systems, you might have just ~1 read per minute, but possibly
	// millions of reads in the same time.
	writeToReadRatio = 1000

	// Delta used in runAdd.
	delta = 0.5
)

func runInc(b *testing.B, c Counter, concurrency int) {
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

func runAdd(b *testing.B, c Counter, concurrency int) {
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

func BenchmarkMutexCounterInc1(b *testing.B) {
	runInc(b, &MutexCounter{}, 1)
}

func BenchmarkMutexCounterAdd1(b *testing.B) {
	runAdd(b, &MutexCounter{}, 1)
}

func BenchmarkRWMutexCounterInc1(b *testing.B) {
	runInc(b, &RWMutexCounter{}, 1)
}

func BenchmarkRWMutexCounterAdd1(b *testing.B) {
	runAdd(b, &RWMutexCounter{}, 1)
}

func BenchmarkAtomicIntCounterInc1(b *testing.B) {
	runInc(b, new(AtomicIntCounter), 1)
}

func BenchmarkAtomicIntCounterAdd1(b *testing.B) {
	runAdd(b, new(AtomicIntCounter), 1)
}

func BenchmarkNaiveCounterInc1(b *testing.B) {
	runInc(b, new(NaiveCounter), 1)
}

func BenchmarkNaiveCounterAdd1(b *testing.B) {
	runAdd(b, new(NaiveCounter), 1)
}

func BenchmarkNaiveIntCounterInc1(b *testing.B) {
	runInc(b, new(NaiveIntCounter), 1)
}

func BenchmarkNaiveIntCounterAdd1(b *testing.B) {
	runAdd(b, new(NaiveIntCounter), 1)
}

func BenchmarkAtomicSpinningCounterInc1(b *testing.B) {
	runInc(b, new(AtomicSpinningCounter), 1)
}

func BenchmarkAtomicSpinningCounterAdd1(b *testing.B) {
	runAdd(b, new(AtomicSpinningCounter), 1)
}

func BenchmarkSyncChannelCounterInc1(b *testing.B) {
	runInc(b, NewChannelCounter(0), 1)
}

func BenchmarkSyncChannelCounterAdd1(b *testing.B) {
	runAdd(b, NewChannelCounter(0), 1)
}

func BenchmarkBufferedChannelCounterInc1(b *testing.B) {
	runInc(b, NewChannelCounter(1024), 1)
}

func BenchmarkBufferedChannelCounterAdd1(b *testing.B) {
	runAdd(b, NewChannelCounter(1024), 1)
}

func BenchmarkMutexCounterInc10(b *testing.B) {
	runInc(b, &MutexCounter{}, 10)
}

func BenchmarkMutexCounterAdd10(b *testing.B) {
	runAdd(b, &MutexCounter{}, 10)
}

func BenchmarkRWMutexCounterInc10(b *testing.B) {
	runInc(b, &RWMutexCounter{}, 10)
}

func BenchmarkRWMutexCounterAdd10(b *testing.B) {
	runAdd(b, &RWMutexCounter{}, 10)
}

func BenchmarkAtomicIntCounterInc10(b *testing.B) {
	runInc(b, new(AtomicIntCounter), 10)
}

func BenchmarkAtomicIntCounterAdd10(b *testing.B) {
	runAdd(b, new(AtomicIntCounter), 10)
}

func BenchmarkNaiveCounterInc10(b *testing.B) {
	runInc(b, new(NaiveCounter), 10)
}

func BenchmarkNaiveCounterAdd10(b *testing.B) {
	runAdd(b, new(NaiveCounter), 10)
}

func BenchmarkNaiveIntCounterInc10(b *testing.B) {
	runInc(b, new(NaiveIntCounter), 10)
}

func BenchmarkNaiveIntCounterAdd10(b *testing.B) {
	runAdd(b, new(NaiveIntCounter), 10)
}

func BenchmarkAtomicSpinningCounterInc10(b *testing.B) {
	runInc(b, new(AtomicSpinningCounter), 10)
}

func BenchmarkAtomicSpinningCounterAdd10(b *testing.B) {
	runAdd(b, new(AtomicSpinningCounter), 10)
}

func BenchmarkSyncChannelCounterInc10(b *testing.B) {
	runInc(b, NewChannelCounter(0), 10)
}

func BenchmarkSyncChannelCounterAdd10(b *testing.B) {
	runAdd(b, NewChannelCounter(0), 10)
}

func BenchmarkBufferedChannelCounterInc10(b *testing.B) {
	runInc(b, NewChannelCounter(1024), 10)
}

func BenchmarkBufferedChannelCounterAdd10(b *testing.B) {
	runAdd(b, NewChannelCounter(1024), 10)
}

func BenchmarkMutexCounterInc100(b *testing.B) {
	runInc(b, &MutexCounter{}, 100)
}

func BenchmarkMutexCounterAdd100(b *testing.B) {
	runAdd(b, &MutexCounter{}, 100)
}

func BenchmarkRWMutexCounterInc100(b *testing.B) {
	runInc(b, &RWMutexCounter{}, 100)
}

func BenchmarkRWMutexCounterAdd100(b *testing.B) {
	runAdd(b, &RWMutexCounter{}, 100)
}

func BenchmarkAtomicIntCounterInc100(b *testing.B) {
	runInc(b, new(AtomicIntCounter), 100)
}

func BenchmarkAtomicIntCounterAdd100(b *testing.B) {
	runAdd(b, new(AtomicIntCounter), 100)
}

func BenchmarkNaiveCounterInc100(b *testing.B) {
	runInc(b, new(NaiveCounter), 100)
}

func BenchmarkNaiveCounterAdd100(b *testing.B) {
	runAdd(b, new(NaiveCounter), 100)
}

func BenchmarkNaiveIntCounterInc100(b *testing.B) {
	runInc(b, new(NaiveIntCounter), 100)
}

func BenchmarkNaiveIntCounterAdd100(b *testing.B) {
	runAdd(b, new(NaiveIntCounter), 100)
}

func BenchmarkAtomicSpinningCounterInc100(b *testing.B) {
	runInc(b, new(AtomicSpinningCounter), 100)
}

func BenchmarkAtomicSpinningCounterAdd100(b *testing.B) {
	runAdd(b, new(AtomicSpinningCounter), 100)
}

func BenchmarkSyncChannelCounterInc100(b *testing.B) {
	runInc(b, NewChannelCounter(0), 100)
}

func BenchmarkSyncChannelCounterAdd100(b *testing.B) {
	runAdd(b, NewChannelCounter(0), 100)
}

func BenchmarkBufferedChannelCounterInc100(b *testing.B) {
	runInc(b, NewChannelCounter(1024), 100)
}

func BenchmarkBufferedChannelCounterAdd100(b *testing.B) {
	runAdd(b, NewChannelCounter(1024), 100)
}
