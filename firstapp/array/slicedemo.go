package main

import (
	"fmt"
)

func main (){
	//slice default pass by reference
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Length of slice a is %v", len(a))
	fmt.Println("Capacity of slice a is %v", cap(a))
}