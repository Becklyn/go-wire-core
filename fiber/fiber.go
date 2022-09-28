package fiber

import (
	"context"

	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/env"
	"github.com/Becklyn/go-wire-core/metrics"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

type MiddlewareHandlerMap map[string][]fiber.Handler

func NewEmptyMiddlewareHandlerMap() *MiddlewareHandlerMap {
	return nil
}

type NewFiberOptions struct {
	Logger     *logrus.Logger
	Middleware *MiddlewareHandlerMap
}

func NewFiber(options *NewFiberOptions) *fiber.App {
	app := fiber.New()

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${latency} - ${status} ${method} ${path}\n",
		Output: options.Logger.Writer(),
	}))

	app.Use(errorMiddleware(options.Logger))
	app.Use(metrics.NewFiberMiddleware())

	if options.Middleware != nil {
		for path, handlers := range *options.Middleware {
			for _, handler := range handlers {
				if handler == nil {
					continue
				}

				if path == "" {
					app.Use(handler)
				} else {
					app.Use(path, handler)
				}
			}
		}
	}

	return app
}

type UseFiberOptions struct {
	Lifecycle *app.Lifecycle
	Fiber     *fiber.App
	Logger    *logrus.Logger
}

func UseFiber(options *UseFiberOptions) {
	addr := env.StringWithDefault("FIBER_ADDR", ":3000")

	options.Lifecycle.OnStart(func(ctx context.Context) error {
		go func() {
			if err := options.Fiber.Listen(addr); err != nil {
				options.Logger.Fatal(err)
			}
		}()

		return nil
	})

	options.Lifecycle.OnStop(func(ctx context.Context) error {
		return options.Fiber.Shutdown()
	})
}
