package core

type Counter struct {
	Value int
}

func (c *Counter) Increment() int {
	c.Value++
	return c.Value
}

func NewCounter(start int) *Counter {
	return &Counter{
		Value: start,
	}
}
