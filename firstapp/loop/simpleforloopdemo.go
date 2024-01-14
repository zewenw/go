package main

import (
	"fmt"
)

func main(){
	for i, j := 0, 0; i < 5; i, j = i + 1, j + 2 {
		fmt.Println(i, j)
	}
	fmt.Println("=================")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	fmt.Println("=================")
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("=================")
	j := 1
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}
	fmt.Println("=================")
	for k := 0; k < 10; k++ {
		if k % 2  == 0 {
			continue
		}
		fmt.Println(k)
	}
}