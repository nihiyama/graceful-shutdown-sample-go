package bl

import (
	"fmt"
	"sync"
	"time"
)

func ProcessData(wg *sync.WaitGroup, workerID int, dataChan chan int) {
	defer wg.Done()
	for {
		data, ok := <-dataChan
		if !ok && len(dataChan) == 0 {
			fmt.Println("nothing!!")
			return
		}
		fmt.Printf("dataChan length: %d, ", len(dataChan))
		fmt.Printf("Worker %d processing data: %v\n", workerID, data)
		time.Sleep(1 * time.Second)
	}
}
