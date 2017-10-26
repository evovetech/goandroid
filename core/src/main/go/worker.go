package core

// +build android

type goWorker struct {
	ScheduleWorker
	*goDisposable
}

func (w *goWorker) Schedule(r Runnable, nanos int) Disposable {
	// TODO:
	return &goDisposable{}
}
