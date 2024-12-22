package counter

import (
	"math/rand"
	"runtime"
	"sync"
)

type ApproxCounter struct {
	v        int64
	nbuckets int
	thresh   int64
	counters []int64

	localMus []*sync.Mutex
	globalMu sync.Mutex
}

func (c *ApproxCounter) Inc() {
	bucket := rand.Int() % c.nbuckets
	c.localMus[bucket].Lock()
	defer c.localMus[bucket].Unlock()
	c.counters[bucket]++
	if c.counters[bucket] > c.thresh {
		c.globalMu.Lock()
		c.v += c.counters[bucket]
		c.globalMu.Unlock()
		c.counters[bucket] = 0
	}
}

func (c *ApproxCounter) Get() int64 {
	c.globalMu.Lock()
	defer c.globalMu.Unlock()
	return c.v
}

func (c *ApproxCounter) Reset() {
	c.globalMu.Lock()
	for i, localMu := range c.localMus {
		localMu.Lock()
		defer localMu.Unlock()
		c.counters[i] = 0
	}
	c.v = 0
	defer c.globalMu.Unlock()
}

// The count is off my at most `runtime.GOMAXPROCS(0) * thresh`
func NewApproxCounter(thresh int64) *ApproxCounter {
	nbuckets := runtime.GOMAXPROCS(0)
	counters := make([]int64, nbuckets)

	localMus := make([]*sync.Mutex, 0, nbuckets)
	for i := 0; i < nbuckets; i++ {
		var mu sync.Mutex
		localMus = append(localMus, &mu)
	}

	return &ApproxCounter{
		v:        0,
		nbuckets: nbuckets,
		thresh:   thresh,
		counters: counters,

		localMus: localMus,
	}
}
