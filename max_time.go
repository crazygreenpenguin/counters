package counters

import "sync/atomic"

type MaxTime int64

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

func (m *MaxTime) Get() int64 {
	return atomic.LoadInt64((*int64)(m))
}

func (m *MaxTime) Reset() {
	atomic.StoreInt64((*int64)(m), 0)
}
