package utils

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown() (ctxSig context.Context, stop context.CancelFunc) {
	ctxSig, stop = signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT,
	)

	go func() {
		<-ctxSig.Done()
		ctxTime, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		for {
			select {
			case <-ctxTime.Done():
				fmt.Println("force shutdown!!!!!")
				cancel()
				os.Exit(1)
			case <-time.After(1 * time.Second):
				fmt.Println("cancel..")
			}
		}
	}()
	return ctxSig, stop
}
