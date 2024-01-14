package main

import (
	"fmt"
)

func main() {

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("The length of Students array is : %v\n", len(students))


	grades1 := [...]int{97, 85, 83}
	fmt.Printf("Grades:  %v\n", grades1)

	grades2 := [3]int{97, 85, 83}
	fmt.Printf("Grades:  %v\n", grades2)

	grade1 := 97
	grade2 := 98
	grade3 := 99

	fmt.Printf("Grades: %v, %v, %v", grade1, grade2, grade3)

}
