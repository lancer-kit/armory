package routines

import (
	"context"
	"fmt"
	"time"
)

// DummyWorker is a simple realization of the Worker interface.
type DummyWorker struct {
	tickDuration time.Duration
	ctx          context.Context
}

// Init returns new instance of the `DummyWorker`.
func (*DummyWorker) Init(parentCtx context.Context) Worker {
	return &DummyWorker{
		ctx:          parentCtx,
		tickDuration: time.Second,
	}
}

// RestartOnFail determines the need to restart the worker, if it stopped
func (s *DummyWorker) RestartOnFail() bool {
	return true
}

// Run start job execution.
func (s *DummyWorker) Run() {
	ticker := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("I'm alive")
		case <-s.ctx.Done():
			ticker.Stop()
			fmt.Println("End job")
			return
		}
	}
}
