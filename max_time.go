package counters

import "sync/atomic"

// MaxTime extend int64 type for thread safe CAS-based primitives
type MaxTime int64

// Store save val in MaxTime if val greater than current MaxTime value
func (m *MaxTime) Store(val int64) {
	for {
		tmpVal := atomic.LoadInt64((*int64)(m))
		if tmpVal < val {
			if atomic.CompareAndSwapInt64((*int64)(m), tmpVal, val) {
				return
			}
		} else {
			if atomic.CompareAndSwapInt64((*int64)(m), tmpVal, tmpVal) {
				return
			}
		}
	}
}

// Get current MaxTime value
func (m *MaxTime) Get() int64 {
	return atomic.LoadInt64((*int64)(m))
}

// Reset set MaxTime value to zero
func (m *MaxTime) Reset() {
	atomic.StoreInt64((*int64)(m), 0)
}
