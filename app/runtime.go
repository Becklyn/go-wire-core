package app

import (
	"context"
	"os"
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	if err := a.lifecycle.Start(ctx); err != nil {
		a.logger.Error(err)
		return
	}

	select {
	case <-ctx.Done():
		a.logger.Info("runtime context done")
	case sig := <-sigChan:
		a.logger.Infof("runtime received signal: %s", sig)
	}

	if err := a.lifecycle.Stop(); err != nil {
		a.logger.Error(err)
		return
	}
	a.logger.Info("runtime stopped")
	os.Exit(0)
}
