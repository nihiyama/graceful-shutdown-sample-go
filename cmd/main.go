package main

import (
	"sync"

	"github.com/nihiyama/graceful-shutdown-sample/internal/bl"
	"github.com/nihiyama/graceful-shutdown-sample/internal/utils"
)

func main() {
	var wg sync.WaitGroup
	ctx, stop := utils.GracefulShutdown()
	defer stop()

	dataChan := make(chan int, 10)

	wg.Add(1)
	go bl.GetData(ctx, &wg, dataChan)

	workerNum := 5
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go bl.ProcessData(&wg, i, dataChan)
	}
	wg.Wait()
}
