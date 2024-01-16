package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// buffer channel
	ch := make(chan int, 50)
	wg.Add(2)

	// receive only channel, consume message from channel
	go func(ch <-chan int) {
		// automatic close a chan
		// for i := range ch {
		// 	fmt.Println(i)
		// }

		// close channel manually
		for {  
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	// send only channel, produce message to channel
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
