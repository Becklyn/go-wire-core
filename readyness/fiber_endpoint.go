package readyness

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UseFiberEndpointOptions struct {
	Fiber     *fiber.App
	Readyness *Service
}

func UseFiberEndpoint(options *UseFiberEndpointOptions) {
	options.Fiber.Get("/ready", options.Handle)
}

func (o *UseFiberEndpointOptions) Handle(c *fiber.Ctx) error {
	if ready, component := o.Readyness.IsReady(); !ready {
		return c.Status(fiber.StatusServiceUnavailable).
			SendString(fmt.Sprintf("Service not ready: uninitialized component: %s", component))
	}
	return c.Status(fiber.StatusOK).SendString("Service ready")
}
