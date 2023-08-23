package fiber

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/env"
	"github.com/Becklyn/go-wire-core/metrics"
	"github.com/bytedance/sonic"
	"github.com/fraym/golog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
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
	jsonEncoder := func() utils.JSONMarshal {
		if env.StringWithDefault("FIBER_JSON_ENCODER", "sonic") == "sonic" {
			return sonic.Marshal
		}
		return json.Marshal
	}()
	jsonDecoder := func() utils.JSONUnmarshal {
		if env.StringWithDefault("FIBER_JSON_DECODER", "sonic") == "sonic" {
			return sonic.Unmarshal
		}
		return json.Unmarshal
	}()

	app := fiber.New(fiber.Config{
		BodyLimit:   env.IntWithDefault("HTTP_REQUEST_BODY_LIMIT", 4) * 1024 * 1024,
		JSONEncoder: jsonEncoder,
		JSONDecoder: jsonDecoder,
	})

	app.Use(fiberlog.New(fiberlog.Config{
		Format: "${latency} - ${status} ${method} ${path}\n",
		Output: options.Logger.Writer(golog.DebugLevel),
	}))

	app.Use(errorMiddleware(options.Logger))
	app.Use(metrics.NewFiberMiddleware())

	corsOrigins := env.StringWithDefault("CORS_ALLOW_ORIGINS", "")
	corsHeaders := env.StringWithDefault("CORS_ALLOW_HEADERS", "")
	corsExposeHeaders := env.StringWithDefault("CORS_EXPOSE_HEADERS", "")
	corsCredentials := env.BoolWithDefault("CORS_ALLOW_CREDENTIALS", false)

	if corsOrigins != "" || corsHeaders != "" || corsExposeHeaders != "" || corsCredentials {
		app.Use(cors.New(cors.Config{
			AllowOrigins:     corsOrigins,
			AllowHeaders:     corsHeaders,
			AllowCredentials: corsCredentials,
			ExposeHeaders:    corsExposeHeaders,
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
