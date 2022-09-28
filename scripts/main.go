package main

import (
	"time"

	"github.com/judegiordano/startup/pkg/logger"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	logger.Init()
	logger.Info("done")
	end := time.Now().UnixNano() / int64(time.Millisecond)
	logger.Info("operation complete in", end-start, "\bms")
}
