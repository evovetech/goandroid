package core

// +build android

type gopool struct {
	Once

	workers []*gopoolWorker
	queue   requestQueue
	done    chan int
}

type gopoolWorker struct {
	in chan *request
}

func (p *gopool) init(size int) {
	p.Once.init(func() {
		p.workers = make([]*gopoolWorker, size)
		p.queue.init(size)
		p.done = make(chan int, 1)
	})
}
