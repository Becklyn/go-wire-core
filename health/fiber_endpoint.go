package health

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UseFiberEndpointOptions struct {
	Fiber  *fiber.App
	Health *Service
}

func UseFiberEndpoint(options *UseFiberEndpointOptions) {
	options.Fiber.Get("/health", options.Handle)
}

func (o *UseFiberEndpointOptions) Handle(c *fiber.Ctx) error {
	if healthy, reason := o.Health.IsHealthy(); !healthy {
		return c.Status(fiber.StatusServiceUnavailable).
			SendString(fmt.Sprintf("Service unavailable: %s", reason))
	}
	return c.Status(fiber.StatusOK).SendString("Service healthy")
}
