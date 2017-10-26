package core

// +build android

import (
	"fmt"
)

type goWorker struct {
	ScheduleWorker
	*goDisposable
}

func (w *goWorker) Schedule(r Runnable, nanos int) (Disposable, error) {
	return w.doSchedule(r, nanos)
}

func (w *goWorker) doSchedule(r Runnable, nanos int) (d Disposable, err error) {
	gd := &goDisposable{}
	d = gd
	fmt.Printf("scheduling in %d nanos\n", nanos)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			switch t := r.(type) {
			case error:
				err = t
			default:
				err = fmt.Errorf("Error: %s\n", t)
			}
		}
	}()
	r.Run()
	gd.Dispose()
	return
}
