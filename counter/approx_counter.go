package counter

type ApproxCounter struct {
	v int64
}

func (c *ApproxCounter) Inc() {}

func (c *ApproxCounter) Get() int64 { return 0 }

func (c *ApproxCounter) Reset() {}
