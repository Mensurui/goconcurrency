package errors

import (
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Deadlock struct {
	hl    hclog.Logger
	syncM *sync.Mutex
}

type value struct {
	mu    sync.Mutex
	value int
}

func NewDeadlock(hl hclog.Logger, syncM *sync.Mutex) *Deadlock {
	return &Deadlock{
		hl:    hl,
		syncM: syncM,
	}
}

func (dl Deadlock) DeadlockVisual() error {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	dl.syncM.Lock()
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
	return nil
}
