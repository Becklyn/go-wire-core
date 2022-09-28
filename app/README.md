# App

Your application runtime basis. Observes SIGINT and SIGTERM and calls lifecycle hooks.

## Lifecycle

Use the `app.Lifecycle` struct to add lifecycle hooks to your application runtime.
Register hooks using the `OnStart` and `OnStop` functions.

## Runtime

Use the `app.Runtime` struct to start the applications lifecycle and listen to SIGINT and SIGTERM system events.

## Usage

Example: a simple application using the fiber webserver.

File: `src/app.go`

```go
package app

import (
	"github.com/Becklyn/go-wire-core/app"
	coreFiber "github.com/Becklyn/go-wire-core/fiber"
)

type AppOptions struct {
	Runtime         *app.Runtime
	UseFiberOptions *coreFiber.UseFiberOptions
}

type App struct {
	AppOptions
}

func newApp(options *AppOptions) *App {
	coreFiber.UseFiber(options.UseFiberOptions)

	return &App{
		*options,
	}
}

func (a *App) Run() {
	a.Runtime.Start()
}
```

File: `src/app/wire.go`

```go
//go:build wireinject

package app

import (
	core "github.com/Becklyn/go-wire-core"
	"github.com/google/wire"
)

func New() *App {
	wire.Build(
		core.Default,
		wire.Struct(new(AppOptions), "*"),
		newApp,
	)
	return nil
}

```

File: `src/main.go`

```go
package main

import (
	"github.com/your-app/src/app"
)

func main() {
	app.New().Run()
}
```
