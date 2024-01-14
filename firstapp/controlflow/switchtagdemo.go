package main

import (
	"fmt"
)

func main(){
	switch i := 2 + 4; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("not one or two")
	}
}