package main

import (
	"sync"
	"time"

	"github.com/judegiordano/startup/pkg/logger"
)

func counter(size int) int {
	var wg sync.WaitGroup
	wg.Add(size)
	count := make(chan int, size)
	defer close(count)
	count <- 0
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * 1_000)
			count <- 1 + <-count
		}()
	}
	wg.Wait()
	return <-count
}

func main() {
	start := time.Now()
	// start
	count := counter(2_000)
	logger.Info(count)
	// done
	time.Sleep(time.Millisecond * 10)
	logger.Info("operation complete in", time.Since(start).Milliseconds(), "\bms")
}
