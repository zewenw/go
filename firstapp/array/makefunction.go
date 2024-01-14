package main

import (
	"fmt"
)

func main(){
	a := make([]int, 3, 100)
	fmt.Printf("Length of slice a is %v\n", len(a))
	fmt.Printf("Capacity of slice a is %v\n", cap(a))
	fmt.Printf("===============\n")

	//remove from the middle, all operation happened on reference
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(d)
	c := append(d[:2], d[3:]...)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("===============\n")

	b := []int{}
	b = append(b, 1)
	fmt.Println(b)
	fmt.Printf("Length of slice a is %v\n", len(b))
	fmt.Printf("Capacity of slice a is %v\n", cap(b))
	b = append(b, 2, 3, 4, 5)
	fmt.Println(b)
	fmt.Printf("Length of slice a is %v\n", len(b))
	fmt.Printf("Capacity of slice a is %v\n", cap(b))
	b = append(b, 2, 3, 4, 5)
	fmt.Println(b)
	fmt.Printf("Length of slice a is %v\n", len(b))
	fmt.Printf("Capacity of slice a is %v", cap(b))
}