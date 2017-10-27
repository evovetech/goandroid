package core

// +build android

import (
	"fmt"
	"runtime"
	"sync"
)

const debugMaxProcs bool = false

var s goScheduler

type goScheduler struct {
	Scheduler

	mu   sync.Mutex
	ctx  context
	pool gopool
}

func scheduler() Scheduler {
	sch := s.init()
	fmt.Printf("scheduler(): &s       = %p\n", &s)
	fmt.Printf("scheduler(): s.init() = %p\n", sch)
	return sch
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

func (s *goScheduler) init() *goScheduler {
	s.Start()
	return s
}

func (s *goScheduler) Start() {
	s.mu.Lock()
	s.pool.init(poolSize())
	s.ctx, _ = rootContextWithCancel()
	s.mu.Unlock()
}

func (s *goScheduler) Shutdown() {
	s.mu.Lock()
	s.ctx.Dispose()
	s.pool = gopool{}
	s.mu.Unlock()
}

func (s *goScheduler) CreateWorker() Worker {
	//TODO:
	fmt.Printf("s = %p\n", s)
	return newWorker(s.ctx)
}

func (s *goScheduler) Schedule(r Runnable, delayNanos int64) (Disposable, error) {
	// TODO:
	fmt.Printf("s = %p\n", s)
	return s.CreateWorker().Schedule(r, delayNanos)
}
