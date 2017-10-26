package core

// +build android

type goDisposable struct {
	Disposable
	p int
}

func (d *goDisposable) Dispose() {
	d.p = 1
}

func (d *goDisposable) IsDisposed() bool {
	return d.p != 0
}
