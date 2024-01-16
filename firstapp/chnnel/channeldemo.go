package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
	ch := make(chan int)
	wg.Add(2)
	// receiver
	go func(){
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()

	// sender
	go func ()  {
		i := 42
		ch <- i
		i = 12
		wg.Done()
	}()

	wg.Wait()
}