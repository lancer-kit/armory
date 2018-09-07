package debugworker

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-kit/routines"
)

type Debugger struct {
	Name       string
	ParentPort int
	ParentIP   string

	Logger *logrus.Entry
	ctx    context.Context
}

func (d *Debugger) Init(parentCtx context.Context) routines.Worker {
	logger, ok := parentCtx.Value(routines.CtxKeyLog).(*logrus.Entry)
	if !ok {
		logger = log.Default
	}

	d.Logger = logger.WithField("worker", "debugger")

	if d.Name == "" {
		d.Name = "debugger"
	}

	d.ctx = parentCtx
	return d
}

func (d *Debugger) RestartOnFail() bool {
	return true
}

func (d *Debugger) Run() {
	go d.runRouter()

	for {
		select {
		case <-d.ctx.Done():
			d.Logger.Info("End job")
			return
		}
	}
}

func (d *Debugger) runRouter() {

}
