package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/judegiordano/startup/api/dev"
	"github.com/judegiordano/startup/internal"
	"github.com/judegiordano/startup/pkg/logger"
)

func main() {
	logger.Init()
	config := internal.LoadConfig()
	app := fiber.New(fiber.Config{
		ErrorHandler: internal.ErrorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(cors.New())
	api := app.Group("/api")
	dev.Router(api)
	err := app.Listen(fmt.Sprintf(":%v", config.Port))
	if err != nil {
		logger.Fatal(err)
	}
}
