package core

// +build android

type Runnable interface {
	Run() error
}

type Disposable interface {
	Dispose()
	IsDisposed() bool
}

type ScheduleWorker interface {
	Schedule(r Runnable, nanos int) (Disposable, error)
}

type Worker interface {
	ScheduleWorker
	Disposable
}

type Scheduler interface {
	ScheduleWorker
	CreateWorker() Worker
}

func GoScheduler() Scheduler {
	return scheduler()
}
