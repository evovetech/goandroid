package core

// +build android

import (
	"Java/java/lang"
)

type Worker interface {
	Schedule(r lang.Runnable, millis int) *GoDisposable
}

type GoWorker struct {
	Worker
	*GoDisposable
}

func (w *GoWorker) Schedule(r lang.Runnable, millis int) *GoDisposable {
	return NewGoDisposable()
}

func NewGoWorker() *GoWorker {
	return &GoWorker{
		GoDisposable: NewGoDisposable(),
	}
}
