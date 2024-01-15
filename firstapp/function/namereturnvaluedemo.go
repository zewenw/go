package main

import (
	"fmt"
)

func main() {
	s := sum("The sum is", 1, 2, 3, 4, 5)
	fmt.Println("The result is ", s)
}

// variadic parameter only allow exists at the end of parameters
func sum(msg string, values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
	return 
}
