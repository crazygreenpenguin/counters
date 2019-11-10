package counters

import "sync/atomic"

type Uint64Counter uint64

func (c *Uint64Counter) Inc() {
	atomic.AddUint64((*uint64)(c), 1)
}

func (c *Uint64Counter) Add(val uint64) {
	atomic.AddUint64((*uint64)(c), val)
}

func (c *Uint64Counter) Set(val uint64) {
	atomic.StoreUint64((*uint64)(c), val)
}

func (c *Uint64Counter) Get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}
