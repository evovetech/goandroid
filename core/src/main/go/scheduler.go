package core

// +build android

import (
	"runtime"
	"fmt"
	"container/list"
	"sync"
)

var once sync.Once
var pool *Pool

type requestChannel = chan []*request

type Pool struct {
	workers []*poolWorker
	queue   *requestQueue
}

func GetPool() *Pool {
	once.Do(func() {
		size := runtime.GOMAXPROCS(0)
		fmt.Printf("GOMAXPROCS: %d\n", size)
		pool = &Pool{
			workers: make([]*poolWorker, size),
			queue: newRequestQueue(0),
		}
	})
	return pool
}

type request struct {
	time int
}

type requestQueue struct {
	*list.List
	in chan []*request
}

func newRequestQueue(size int) *requestQueue {
	return &requestQueue{
		List: list.New(),
		in: make(chan []*request, size),
	}
}

type poolWorker struct {
	in chan []*request
}


