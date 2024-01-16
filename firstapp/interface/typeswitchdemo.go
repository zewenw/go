package main

import (
	"fmt"
)

func main(){
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is a int type")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}