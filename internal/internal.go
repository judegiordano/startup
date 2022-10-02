package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/judegiordano/startup/pkg/logger"
)

type Config struct {
	Stage    string
	MongoUri string
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func LoadInt(key string) int {
	s, found := os.LookupEnv(key)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", s))
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		logger.Fatal(err)
	}
	return n
}

func LoadString(key string) string {
	s, found := os.LookupEnv(key)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", s))
	}
	return s
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf("error loading .env: %v", err))
	}
	return Config{
		Stage:    LoadString("STAGE"),
		MongoUri: LoadString("MONGO_URI"),
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()
	logger.Error(err)
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}
	return ctx.Status(code).JSON(ErrorResponse{
		Code:  code,
		Error: message,
	})
}
