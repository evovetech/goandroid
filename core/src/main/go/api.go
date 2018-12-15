package core

// +build android

type Counter interface {
	Increment() int32
	GetValue() int32
}

type Runnable interface {
	Run() error
}

type Disposable interface {
	Dispose()
	IsDisposed() bool
}

type ScheduleWorker interface {
	Schedule(r Runnable, delayNanos int64) (Disposable, error)
}

type Worker interface {
	ScheduleWorker
	Disposable
}

type Scheduler interface {
	ScheduleWorker
	Start()
	Shutdown()
	CreateWorker() Worker
}

func GoScheduler() Scheduler {
	return scheduler()
}
