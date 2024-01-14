package main

import (
	"fmt"
)

func main (){
	var myByte byte

	myByte = 'A'

	fmt.Printf("%v, %T\n", myByte, myByte)

	myString := "hello world"

	for i := 0; i < len(myString); i++ {
		fmt.Printf("%c -> %v\n", myString[i], myString[i])
	}
}