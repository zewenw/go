package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){
	ch := make(chan int)
	wg.Add(2)
	/*
	common channel can buffer data
	which means that if no consumer consume data from it
	then producer will blocked until consumer consume the data in it
	*/ 
	// receive only channel, consume message from channel
	go func(ch <- chan int){
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	// send only channel, produce message to channel
	go func (ch chan <- int){
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}