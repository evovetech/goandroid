package core

// +build android

import (
	ctx "context"
	"time"
)

type context interface {
	ctx.Context
	Disposable
}

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.Background().Deadline()
}

func (*emptyCtx) Done() <-chan struct{} {
	return ctx.Background().Done()
}

func (*emptyCtx) Err() error {
	return ctx.Background().Err()
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return ctx.Background().Value(key)
}

func (*emptyCtx) Dispose() {
	// do nothing
}

func (*emptyCtx) IsDisposed() bool {
	return false
}

var rootCxt = new(emptyCtx)

func rootContext() context {
	return rootCxt
}

func rootContextWithCancel() (context, ctx.CancelFunc) {
	return contextWithCancel(rootContext())
}

func contextWithCancel(parent context) (context, ctx.CancelFunc) {
	c := new(cancelCtx)
	c.Context, c.cancel = ctx.WithCancel(realContext(parent))
	return c, c.cancel
}

func realContext(ctx context) ctx.Context {
	for {
		switch c := ctx.(type) {
		case *cancelCtx:
			return c.Context
		default:
			return ctx
		}
	}
}

type cancelCtx struct {
	ctx.Context

	cancel ctx.CancelFunc
}

func (c *cancelCtx) Dispose() {
	c.cancel()
}

func (c *cancelCtx) IsDisposed() bool {
	return c.Err() != nil
}
