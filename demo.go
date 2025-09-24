package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-signalChan:
			fmt.Println("going to stop")
			os.Exit(0)
		case <-ticker.C:
			fmt.Println("working")
			time.Sleep(3 * time.Second)
		}
	}
}
