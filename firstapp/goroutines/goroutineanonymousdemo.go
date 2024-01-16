package main

import (
	"fmt"
	"sync"
	// "time"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "hello"
	wg.Add(1)
	go func() {
		fmt.Println("no argument", msg)
		wg.Done()
	}()

	msg = "Good Bye"
	// time.Sleep(time.Millisecond)
	wg.Wait()
}
