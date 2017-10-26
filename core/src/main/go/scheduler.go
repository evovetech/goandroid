package core

// +build android

import (
	"fmt"
	"runtime"
	"sync"
)

var once sync.Once
var scheduler *goScheduler

type goScheduler struct {
	Scheduler
	pool *pool
}

func (s *goScheduler) init(size int) *goScheduler {
	s.pool = new(pool).
		init(size)
	return s
}

func (s *goScheduler) CreateWorker() Worker {
	//TODO:
	return &goWorker{}
}

func (s *goScheduler) Schedule(r Runnable, nanos int) (Disposable, error) {
	// TODO:
	return s.CreateWorker().
		Schedule(r, nanos)
}

func GoScheduler() Scheduler {
	once.Do(func() {
		size := runtime.GOMAXPROCS(0)
		fmt.Printf("GOMAXPROCS: %d\n", size)
		scheduler = new(goScheduler).
			init(size)
	})
	return scheduler
}
