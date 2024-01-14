package main

import (
	"fmt"
)

func main (){
	//follow is an array which there are three dots
	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:] //slice of all elements
	c := a[3:] // slice from 3th element to the end
	d := a[:6]	// slice first 6th element
	e := a[3:6] // slice the 4th, 5th, 6th elements
	//for slice function, the first argument is inclusive, the second argument is exclusive
	a[4] = 53
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}