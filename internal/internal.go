package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/judegiordano/startup/pkg/logger"
)

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Config struct {
	Port int
}

func LoadString(key string) string {
	s, found := os.LookupEnv(key)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", s))
	}
	return s
}

func LoadInt(key string) int {
	s, found := os.LookupEnv(key)
	if !found {
		logger.Fatal(fmt.Sprintf(".env %v not set", s))
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		logger.Fatal(fmt.Sprintf("error parsing %v", s))
	}
	return n
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err)
	}
	return Config{
		Port: LoadInt("PORT"),
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
