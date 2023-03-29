package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fraym/golog"
)

type Runtime struct {
	lifecycle *Lifecycle
	logger    golog.Logger
}

func NewRuntime(lifecycle *Lifecycle, logger golog.Logger) *Runtime {
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
		a.logger.Error().WithError(err).Write()
		return
	}

	select {
	case <-ctx.Done():
		a.logger.Info().Write("runtime context done")
	case sig := <-sigChan:
		a.logger.Info().Writef("runtime received signal: %s", sig)
	}

	if err := a.lifecycle.Stop(); err != nil {
		a.logger.Error().WithError(err).Write()
		return
	}
	a.logger.Info().Write("runtime stopped")
	os.Exit(0)
}
