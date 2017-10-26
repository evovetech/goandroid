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
	*list.List
	in chan *request
}

func (q *requestQueue) init(size int) *requestQueue {
	q.List = list.New()
	q.in = make(chan *request, size)
	return q
}
