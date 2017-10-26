package core

// +build android

import (
	"fmt"
	"runtime"
)

var s goScheduler

type goScheduler struct {
	Scheduler
	Once

	pool pool
}

func scheduler() Scheduler {
	s.init(func() {
		size := runtime.GOMAXPROCS(0)
		fmt.Printf("GOMAXPROCS: %d\n", size)
		s.pool.init(size)
	})
	fmt.Printf("s = %p\n", &s)
	return &s
}

func (s *goScheduler) CreateWorker() Worker {
	//TODO:
	fmt.Printf("s = %p\n", s)
	return &goWorker{}
}

func (s *goScheduler) Schedule(r Runnable, nanos int) (Disposable, error) {
	// TODO:
	fmt.Printf("s = %p\n", s)
	return s.CreateWorker().Schedule(r, nanos)
}
