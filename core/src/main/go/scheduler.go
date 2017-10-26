package core

// +build android

import (
	"container/list"
	"fmt"
	"runtime"
	"sync"
)

var once sync.Once
var pool *Pool

type Pool struct {
	workers []*poolWorker
	queue   *requestQueue
	done    chan int
}

type poolWorker struct {
	in chan *request
}

func GetPool() *Pool {
	once.Do(func() {
		size := runtime.GOMAXPROCS(0)
		fmt.Printf("GOMAXPROCS: %d\n", size)
		pool = &Pool{
			workers: make([]*poolWorker, size),
			queue:   new(requestQueue).init(0),
			done:    make(chan int, 1),
		}
	})
	return pool
}

type request struct {
	atTime   int
	run      Runnable
	disposed chan int
}

type requestQueue struct {
	*list.List
	in chan *request
}

func (q *requestQueue) init(size int) *requestQueue {
	q.List = list.New()
	q.in = make(chan *request, size)
	return q
}
