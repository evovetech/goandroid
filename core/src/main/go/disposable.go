package core

// +build android

import (
	rxd "Java/io/reactivex/disposables"
)

type GoDisposable struct {
	rxd.Disposable
	p int
}

func NewGoDisposable() *GoDisposable {
	return &GoDisposable{}
}

func (d *GoDisposable) Dispose() {
	d.p = 1
}

func (d *GoDisposable) IsDisposed() bool {
	return d.p != 0
}
