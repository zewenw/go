package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("Thread %v\n", runtime.GOMAXPROCS(-1))
}