package core

// +build android

import (
	"fmt"
	"runtime"
)

const debugMaxProcs bool = false

var s goScheduler

type goScheduler struct {
	Scheduler
	Once

	pool gopool
}

func scheduler() Scheduler {
	s.init(func() {
		s.pool.init(poolSize())
	})
	fmt.Printf("scheduler() = %p\n", &s)
	return &s
}

func poolSize() int {
	var size int
	if debugMaxProcs {
		size = 2
	} else {
		max := runtime.GOMAXPROCS(0)
		size = max*2 + 1
	}
	fmt.Printf("GOMAXPROCS: %d\n", size)
	return size
}

func (s *goScheduler) CreateWorker() Worker {
	//TODO:
	fmt.Printf("s = %p\n", s)
	return &goWorker{}
}

func (s *goScheduler) Schedule(r Runnable, nanos int64) (Disposable, error) {
	// TODO:
	fmt.Printf("s = %p\n", s)
	return s.CreateWorker().Schedule(r, nanos)
}
