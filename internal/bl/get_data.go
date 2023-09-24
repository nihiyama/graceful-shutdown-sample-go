package bl

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

func GetData(ctx context.Context, wg *sync.WaitGroup, dataChan chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			close(dataChan)
			return
		default:
			data := make([]int, 10)
			// dummy get data
			for i := range data {
				data[i] = rand.Intn(100)
			}

			// yield data
			for _, e := range data {
				fmt.Println("get data")
				dataChan <- e
			}
		}
	}
}
