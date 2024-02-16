package app

import "context"

type LifecycleHook func(ctx context.Context) error

type Lifecycle struct {
	onStart    []LifecycleHook
	onStop     []LifecycleHook
	onStopLast []LifecycleHook
}

func NewLifecycle() *Lifecycle {
	return &Lifecycle{}
}

func (l *Lifecycle) Start(ctx context.Context) error {
	for _, hook := range l.onStart {
		if err := hook(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *Lifecycle) Stop() error {
	ctx := context.Background()

	for _, hook := range l.onStop {
		if err := hook(ctx); err != nil {
			return err
		}
	}

	for _, hook := range l.onStopLast {
		if err := hook(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *Lifecycle) OnStart(hook LifecycleHook) {
	l.onStart = append(l.onStart, hook)
}

func (l *Lifecycle) OnStop(hook LifecycleHook) {
	l.onStop = append(l.onStop, hook)
}

func (l *Lifecycle) OnStopLast(hook LifecycleHook) {
	l.onStopLast = append(l.onStopLast, hook)
}
