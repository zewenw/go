package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	//defer keyword will defer execution of code before the end of the method
	defer fmt.Println("middle")
	defer fmt.Println("end")
	defer fmt.Println("end2")

	a := "a - start"
	defer fmt.Println(a)
	a = "a - end"
}
