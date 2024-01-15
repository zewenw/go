package main

import (
	"fmt"
)

func main (){
	var i int = 42
	// here *int means that j is a pointer point to a int reference
	var j *int = &i
	fmt.Println(i, j)
	fmt.Println(&i, j)
	fmt.Println(i, *j)
	i = 27
	fmt.Println(i, *j)
	*j = 58
	fmt.Println(i, *j)


	fmt.Println("==================")
	a := 42
	// b get a copy of a instead of reference
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
}