package counter

import "sync"

type LockCounter struct {
	mu sync.Mutex
	v  int64
}

func (c *LockCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v++
}

func (c *LockCounter) Get() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func (c *LockCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v = 0
}

func NewLockCounter() *LockCounter {
	return &LockCounter{v: 0}
}
