package main

import (
	"fmt"
)

func main(){
	sum("The sum is", 1, 2, 3, 4, 5)
}

// variadic parameter only allow exists at the end of parameters
func sum(msg string, values ...int){
	fmt.Println(values)
	result := 0
	for _, v := range values{
		result += v
	}
	fmt.Println(msg, result)
}
