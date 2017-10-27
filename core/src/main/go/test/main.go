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
	var fin int
	count := 5
	done := make(chan int)
	doneFunc := func() {
		done <- 1
	}
	scheduler := core.GoScheduler()
	r := func(num int) RunnableFunc {
		return func() error {
			defer doneFunc()
			fmt.Printf("run #%d\n", num)
			return nil
		}
	}
	for i := 0; i < count; i++ {
		d, _ := scheduler.Schedule(r(i+1), 0)
		if i%2 == 0 {
			d.Dispose()
			fin++
		}
	}

	for fin < count {
		select {
		case c := <-done:
			fin += c
		}
	}
}
