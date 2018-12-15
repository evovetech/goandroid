package core

// +build android

type counter struct {
	value int32
}

func (c *counter) Increment() int32 {
	c.value++
	return c.value
}

func (c *counter) GetValue() int32 {
	return c.value
}

func NewCounter(start int32) Counter {
	return &counter{
		value: start,
	}
}
