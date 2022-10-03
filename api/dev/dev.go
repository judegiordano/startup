package dev

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/judegiordano/startup/pkg/mongo"
)

type Health struct {
	Message string `json:"message"`
}

func healthCheck(c *fiber.Ctx) error {
	c.Status(418)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	temp := mongo.Collection("temp")
	doc := Health{Message: "☕"}
	if _, err := temp.InsertOne(ctx, &doc); err != nil {
		return fiber.NewError(500, err.Error())
	}
	return c.JSON(&doc)
}

func Router(r fiber.Router) {
	handler := r.Group("/dev")
	handler.Get("/ping", healthCheck)
}
