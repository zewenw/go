package main

import (
	"fmt"
)

func main(){
	/*
	start
	this was deferred	
	panic: something bad happened
	*/
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
}