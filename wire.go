//go:build wireinject

package core

import (
	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/env"
	"github.com/Becklyn/go-wire-core/logging"
	"github.com/google/wire"
)

var Logger = wire.NewSet(
	env.New,
	logging.New,
)

var App = wire.NewSet(
	app.NewLifecycle,
	app.NewRuntime,
)
