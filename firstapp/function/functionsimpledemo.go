package main

import (
	"fmt"
)

func main(){
	// sayMessage("Hello World!")
	// for i := 0; i < 5; i++ {
	// 	sayMessageone("Hello, Go!", i)
	// }
	sayGreeting("hello", "stacey")
}

func sayGreeting(greeting, name string){
	fmt.Println(greeting, name)
}


func sayMessageone(msg string, idx int){
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}


func sayMessage(msg string){
	fmt.Println(msg)
}