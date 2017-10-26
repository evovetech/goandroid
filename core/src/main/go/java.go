package core

// +build android

import (
	rxd "Java/io/reactivex/disposables"
	"Java/java/lang"
)

type Runnable = lang.Runnable
type Disposable = rxd.Disposable
type ScheduleWorker interface {
	Schedule(r Runnable, nanos int) Disposable
}
type Worker interface {
	ScheduleWorker
	rxd.Disposable
}
type Scheduler interface {
	ScheduleWorker
	CreateWorker() Worker
}
