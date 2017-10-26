package core

// +build android

import (
	rxd "Java/io/reactivex/disposables"
)

type workScheduler interface {
	Schedule(r Runnable, millis int) Disposable
}

type Worker interface {
	rxd.Disposable
	workScheduler
}

type goWorker struct {
	*goDisposable
	workScheduler
}

func (w *goWorker) Schedule(r Runnable, millis int) Disposable {
	// TODO:
	return &goDisposable{}
}

func NewWorker() Worker {
	return &goWorker{
		goDisposable: &goDisposable{},
	}
}
