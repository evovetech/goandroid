package core

// +build android

type pool struct {
	Once

	workers []*poolWorker
	queue   requestQueue
	done    chan int
}

type poolWorker struct {
	in chan *request
}

func (p *pool) init(size int) *pool {
	p.Once.init(func() {
		p.workers = make([]*poolWorker, size)
		p.queue.init(size)
		p.done = make(chan int, 1)
	})
	return p
}
