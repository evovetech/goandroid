package core

// +build android

import (
	"sync"
)

type once struct {
	once sync.Once
}

func (o *once) init(f func()) {
	o.once.Do(f)
}
