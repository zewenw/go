package main

import (
	"fmt"
)

func main(){
	a := [...]int{1, 2, 3}
	//pass by reference
	b := &a
	//deep copy rather than shallow copy
	// b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}