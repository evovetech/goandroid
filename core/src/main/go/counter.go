package core

// +build android

type Counter struct {
	Value int32
}

func (c *Counter) Increment() int32 {
	c.Value++
	return c.Value
}

func NewCounter(start int32) *Counter {
	return &Counter{
		Value: start,
	}
}
