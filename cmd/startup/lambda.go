package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/judegiordano/startup/internal/server"
	"github.com/judegiordano/startup/pkg/logger"
)

func main() {
	logger.Init()
	lambda.Start(server.Handler)
}
