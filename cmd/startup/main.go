package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/judegiordano/startup/pkg/logger"
)

func main() {
	logger.Init()
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(cors.New())
	api := app.Group("/api")
	api.Get("/dev", healthCheck)
	app.Listen(fmt.Sprintf(":%v", 1234))
}

func healthCheck(c *fiber.Ctx) error {
	type health struct {
		Message string `json:"message"`
	}
	c.Status(418)
	return c.JSON(health{
		Message: "☕",
	})
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()
	logger.Error(err)
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}
	ctx.Status(code).JSON(ErrorResponse{
		Code:  code,
		Error: message,
	})
	return nil
}
