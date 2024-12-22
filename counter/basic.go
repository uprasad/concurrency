package counter

type BasicCounter struct {
	v int64
}

func (c *BasicCounter) Inc()       { c.v++ }
func (c *BasicCounter) Get() int64 { return c.v }
func (c *BasicCounter) Reset()     { c.v = 0 }

func NewBasicCounter() *BasicCounter {
	return &BasicCounter{v: 0}
}
