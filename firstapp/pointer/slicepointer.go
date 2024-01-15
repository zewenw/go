package main

import (
	"fmt"
)

func main(){
	/*
	primitive type and struct will passing by value, which means
	when a new variable point to original variable, go will deep copy
	the whole value unless you specify you want to use pointer
	*/

	m := map[string]string{"foo" : "bar", "baz" : "buz"}
	n := m
	fmt.Println(m, n)
	m["foo"] = "qux"
	fmt.Println(m, n)
	fmt.Println("====================")

	/*
	following code initialize a slice which is a pointer point to an array, and b get a copy of pointer a
	this behavior is same with map which is also a pointer point to data
	*/
	i := []int{1, 2, 3}
	j := i
	fmt.Println(i, j)
	i[1] = 42
	fmt.Println(i, j)
	
	fmt.Println("====================")

	// following code initialize a array, and b get a copy of array a
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
}