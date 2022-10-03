package dev

import (
	"fmt"

	"github.com/judegiordano/startup/internal/models/temp"

	"github.com/gofiber/fiber/v2"
)

type Health struct {
	Message string `json:"message"`
}

func healthCheck(c *fiber.Ctx) error {
	c.Status(418)
	doc := temp.Document{Message: "☕"}
	done, err := doc.Save()
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	return c.JSON(Health{Message: fmt.Sprintf("inserted: %v", done.InsertedID)})
}

func Router(r fiber.Router) {
	handler := r.Group("/dev")
	handler.Get("/ping", healthCheck)
}
