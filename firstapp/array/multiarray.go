package main

import (
	"fmt"
)

func main (){

	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{1, 1, 0}
	matrix[2] = [3]int{1, 1, 1}
	fmt.Println(matrix)

	fmt.Printf("===========================\n")
	var identityMatrix [3][3]int = [3][3]int{ [3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)


}