package core

// +build android

type goWorker struct {
	context
}

func newWorker(parent context) *goWorker {
	ctx, _ := contextWithCancel(parent)
	return &goWorker{context: ctx}
}

func (w *goWorker) Schedule(r Runnable, delayNanos int64) (Disposable, error) {
	return w.doSchedule(r, delayNanos)
}

func (w *goWorker) doSchedule(r Runnable, delayNanos int64) (d Disposable, err error) {
	ctx, _ := contextWithCancel(w.context)
	go func() {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			r.Run()
		}
	}()
	d = ctx
	return
}
