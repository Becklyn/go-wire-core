package middleware

import (
	"github.com/Becklyn/go-wire-core/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NoCors() fiber.Handler {
	if env.IsDevelopment() {
		return cors.New()
	}

	return nil
}
