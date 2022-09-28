package main

import (
	"github.com/judegiordano/startup/pkg/logger"
)

func main() {
	logger.Init()
	logger.Info("sup")
}
