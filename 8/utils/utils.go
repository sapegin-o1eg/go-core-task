package utils

import (
	"sync"
)

type CustomWaitGroup struct {
	counter int64
	sem     chan struct{}
	mu      sync.Mutex
	closed  bool
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		sem:    make(chan struct{}, 1),
		closed: false,
	}
}

func (wg *CustomWaitGroup) Add(delta int64) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter += delta

	if wg.counter < 0 {
		panic("CustomWaitGroup: negative counter")
	}

	if wg.counter == 0 && !wg.closed {
		close(wg.sem)
		wg.closed = true
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	wg.mu.Lock()

	if wg.counter == 0 {
		wg.mu.Unlock()
		return
	}

	sem := wg.sem
	wg.mu.Unlock()

	<-sem
}

func (wg *CustomWaitGroup) Reset() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.counter = 0
	wg.closed = false
	wg.sem = make(chan struct{}, 1)
}
