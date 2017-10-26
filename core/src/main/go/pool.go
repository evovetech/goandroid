package core

// +build android

type pool struct {
	workers []*poolWorker
	queue   *requestQueue
	done    chan int
}

type poolWorker struct {
	in chan *request
}

func (p *pool) init(size int) *pool {
	p.workers = make([]*poolWorker, size)
	p.queue = new(requestQueue).
		init(0)
	p.done = make(chan int, 1)
	return p
}
