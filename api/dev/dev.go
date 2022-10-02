package dev

import "github.com/gofiber/fiber/v2"

type health struct {
	Message string `json:"message"`
}

func healthCheck(c *fiber.Ctx) error {
	c.Status(418)
	return c.JSON(health{
		Message: "☕",
	})
}

func Router(r fiber.Router) {
	handler := r.Group("/dev")
	handler.Get("/ping", healthCheck)
}
