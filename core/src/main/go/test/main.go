package main

import (
	"fmt"
	"github.com/evovetech/goandroid/core/src/main/go"
)

type RunnableFunc func() error

func (f RunnableFunc) Run() error {
	return f()
}

func main() {
	count := 5
	scheduler := core.GoScheduler()
	for i := 0; i < count; i++ {
		scheduler.Schedule(RunnableFunc(func() error {
			fmt.Printf("run #%d\n", i)
			return nil
		}), 0)
	}
}
