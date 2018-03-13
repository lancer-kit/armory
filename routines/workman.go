package routines

import (
	"context"
	"fmt"
	"time"
)

// Workman is an interface for async workers
// which launches and manages by the `Chief`.
type Workman interface {
	// Init initializes new instance of the `Workman` implementation.
	Init(context.Context) Workman
	// Run starts the `Workman` instance execution.
	Run()
}

// DummyWorkman is a simple realization of the Workman interface.
type DummyWorkman struct {
	tickDuration time.Duration
	ctx          context.Context
}

// Init returns new instance of the `DummyWorkman`.
func (*DummyWorkman) Init(parentCtx context.Context) Workman {
	return &DummyWorkman{
		ctx:          parentCtx,
		tickDuration: time.Second,
	}
}

// Run start job execution.
func (s *DummyWorkman) Run() {
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
