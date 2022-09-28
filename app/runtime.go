package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Runtime struct {
	lifecycle *Lifecycle
	logger    *logrus.Logger
}

func NewRuntime(lifecycle *Lifecycle, logger *logrus.Logger) *Runtime {
	return &Runtime{
		lifecycle: lifecycle,
		logger:    logger,
	}
}

func (a *Runtime) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := a.lifecycle.Start(ctx); err != nil {
		a.logger.Error(err)
		return
	}

	<-ctx.Done()

	if err := a.lifecycle.Stop(); err != nil {
		a.logger.Error(err)
		return
	}
}
