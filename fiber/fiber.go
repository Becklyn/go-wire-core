package fiber

import (
	"context"
	"time"

	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/env"
	"github.com/Becklyn/go-wire-core/metrics"
	"github.com/fraym/golog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
)

type MiddlewareHandlerMap map[string][]fiber.Handler

func NewEmptyMiddlewareHandlerMap() *MiddlewareHandlerMap {
	return nil
}

type NewFiberOptions struct {
	Logger     golog.Logger
	Middleware *MiddlewareHandlerMap
}

func NewFiber(options *NewFiberOptions) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: env.IntWithDefault("HTTP_REQUEST_BODY_LIMIT", 4) * 1024 * 1024,
	})

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${latency} - ${status} ${method} ${path}\n",
		Output: options.Logger.Writer(golog.DebugLevel),
	}))

	app.Use(errorMiddleware(options.Logger))
	app.Use(metrics.NewFiberMiddleware())

	if corsOrigins := env.StringWithDefault("CORS_ALLOW_ORIGINS", ""); corsOrigins != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: corsOrigins,
		}))
	}

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
	Logger    golog.Logger
}

func UseFiber(options *UseFiberOptions) {
	addr := env.StringWithDefault("FIBER_ADDR", ":3000")

	options.Lifecycle.OnStart(func(ctx context.Context) error {
		go func() {
			if err := options.Fiber.Listen(addr); err != nil {
				options.Logger.Fatal().WithError(err).Write()
			}
		}()

		return nil
	})

	options.Lifecycle.OnStop(func(ctx context.Context) error {
		return options.Fiber.ShutdownWithTimeout(3 * time.Second)
	})
}
