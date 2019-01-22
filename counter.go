package counters

import (
	"sync"
	"time"
)

// Counter - object realize simple counter thread-safe primitive for count stats in other program
type Counter struct {
	lock             sync.Mutex
	counter          uint64
	getRPSLastAccess int64
	getRPSLastValue  uint64
}

// NewCounter return pointer on new counter object with zero value
func NewCounter() *Counter {
	return &Counter{
		counter:          0,
		getRPSLastAccess: time.Now().UnixNano(),
		getRPSLastValue:  0,
	}
}

// Set set counter in value
func (c *Counter) Set(value uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.counter = value
	c.getRPSLastAccess = time.Now().UnixNano()
	c.getRPSLastValue = value
}

// Get return current value of counter without counter modify
func (c *Counter) Get() uint64 {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.counter
}

// GetWithReset return current value of counter and set counter to zero
func (c *Counter) GetWithReset() uint64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	result := c.counter
	c.counter = 0
	c.getRPSLastAccess = time.Now().UnixNano()
	return result
}

// Inc add 1 to counter
func (c *Counter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.counter++
}

// GetRPS returns average counter value diff per second
func (c *Counter) GetRPS() float64 {
	c.lock.Lock()
	defer c.lock.Unlock()

	now := time.Now().UnixNano()
	delta := c.counter - c.getRPSLastValue
	if delta == 0 {
		c.getRPSLastAccess = now
		return 0
	}
	rps := float64(1000000000*delta) / float64(now-c.getRPSLastAccess)
	c.getRPSLastValue = c.counter
	c.getRPSLastAccess = now
	return rps
}
