package debugworker

import (
	"context"

	"gitlab.inn4science.com/gophers/service-kit/routines"
)

type Debugger struct {
}

func (d *Debugger) Init(context.Context) routines.Worker {
	return d
}

func (d *Debugger) RestartOnFail() bool {
	return true
}

func (d *Debugger) Run() {

}
