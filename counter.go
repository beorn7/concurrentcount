package concurrentcount

import (
	"math"
	"sync"
	"sync/atomic"
)

// Counter is the interface of all counter implementations. Prometheus always
// uses float64 as values, so we are doing so here, although a pure integer
// counter would be sufficient (or even more suitable) in most situations.
type Counter interface {
	Get() float64
	Inc()
	Add(float64)
}

type MutexCounter struct {
	mtx   sync.Mutex
	value float64
}

func (c *MutexCounter) Get() float64 {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.value
}

func (c *MutexCounter) Inc() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.value++
}

func (c *MutexCounter) Add(delta float64) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.value += delta
}

type RWMutexCounter struct {
	mtx   sync.RWMutex
	value float64
}

func (c *RWMutexCounter) Get() float64 {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.value
}

func (c *RWMutexCounter) Inc() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.value++
}

func (c *RWMutexCounter) Add(delta float64) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.value += delta
}

// AtomicIntCounter uses an int64 internally.
type AtomicIntCounter int64

func (c *AtomicIntCounter) Get() float64 {
	return float64(atomic.LoadInt64((*int64)(c)))
}

func (c *AtomicIntCounter) Inc() {
	atomic.AddInt64((*int64)(c), 1)
}

// Add ignores the non-integer part of delta.
func (c *AtomicIntCounter) Add(delta float64) {
	atomic.AddInt64((*int64)(c), int64(delta))
}

// NaiveCounter does not apply any locking.
type NaiveCounter float64

func (c *NaiveCounter) Get() float64 {
	return float64(*c)
}

func (c *NaiveCounter) Inc() {
	(*c)++
}

// Add ignores the non-integer part of delta.
func (c *NaiveCounter) Add(delta float64) {
	*c += NaiveCounter(delta)
}

// NaiveIntCounter uses an int64 internally and does not apply any locking.
type NaiveIntCounter int64

func (c *NaiveIntCounter) Get() float64 {
	return float64(*c)
}

func (c *NaiveIntCounter) Inc() {
	(*c)++
}

// Add ignores the non-integer part of delta.
func (c *NaiveIntCounter) Add(delta float64) {
	*c += NaiveIntCounter(delta)
}

// AtomicCounter uses an uint64 internally and still deals with real
// float64's. However, it has to use CompareAndSwapUint64 for that, which might
// fail during contention. It might need to try several times, spinning...
type AtomicSpinningCounter uint64

func (c *AtomicSpinningCounter) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(c)))
}

func (c *AtomicSpinningCounter) Inc() {
	c.Add(1)
}

func (c *AtomicSpinningCounter) Add(delta float64) {
	for {
		oldBits := atomic.LoadUint64((*uint64)(c))
		newBits := math.Float64bits(math.Float64frombits(oldBits) + delta)
		if atomic.CompareAndSwapUint64((*uint64)(c), oldBits, newBits) {
			return
		}
	}
}

type ChannelCounter struct {
	queue chan float64
	value chan float64
}

func NewChannelCounter(bufsize int) Counter {
	c := &ChannelCounter{
		make(chan float64, bufsize),
		make(chan float64),
	}
	go c.loop()
	return c
}

func (c *ChannelCounter) Get() float64 {
	return <-c.value
}

func (c *ChannelCounter) Inc() {
	c.queue <- 1
}

func (c *ChannelCounter) Add(delta float64) {
	c.queue <- delta
}

func (c *ChannelCounter) loop() {
	// In a real-life counter, there needed to be a way to shut it down
	// again.

	var value float64
	for {
		// Outer select: First process all the waiting samples.
		// This is not ideal as it might let Get block forever.
		select {
		case v := <-c.queue:
			value += v
		default:
			select {
			case v := <-c.queue:
				value += v
			case c.value <- value:
				// Do nothing.
			}
		}
	}
}
