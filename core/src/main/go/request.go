package core

// +build android

import (
	"container/list"
)

type request struct {
	atTime   int
	run      Runnable
	disposed chan int
}

type requestQueue struct {
	once
	list.List

	in chan *request
}

func (q *requestQueue) init(size int) *requestQueue {
	q.once.init(func() {
		q.List.Init()
		q.in = make(chan *request, size)
	})
	return q
}
