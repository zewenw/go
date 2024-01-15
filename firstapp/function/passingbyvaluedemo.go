package main

import (
	"fmt"
)

func main(){
	greeting := "hello"
	name := "Stacey"
	passingByValue(greeting, name)
	fmt.Println(name)

	fmt.Println("=====================")
	passingByReference(&greeting, &name)
	fmt.Println(greeting, name)
}

func passingByReference(greeting, name *string){
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func passingByValue(greeting, name string){
	fmt.Println(greeting, name)
	name = "ted"
	fmt.Println(name)
}
