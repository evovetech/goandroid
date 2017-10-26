package core

// +build android

import (
	"fmt"
	"sync"
)

const debug = true

type Initializer interface {
	init()
}

type Once struct {
	Initializer

	once sync.Once
}

func (o *Once) init(f func()) {
	o.once.Do(f)
	if debug {
		fmt.Printf("once = %p\n", o)
	}
}
