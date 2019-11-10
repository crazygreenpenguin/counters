package counters

import "sync/atomic"

type Sequence uint64

func (s *Sequence) Set(val uint64) {
	atomic.StoreUint64((*uint64)(s), val)
}

func (s *Sequence) Get() uint64 {
	return atomic.LoadUint64((*uint64)(s))
}

func (s *Sequence) Next() uint64 {
	for {
		val := atomic.LoadUint64((*uint64)(s))
		if atomic.CompareAndSwapUint64((*uint64)(s), val, val+1) {
			return val + 1
		}
	}
}
