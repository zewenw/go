package main

import (
	"fmt"
)

func main() {
	str := "Hello Go!"
	for k, v := range str {
		fmt.Println(k, string(v))
	}
	fmt.Println("=====================")


	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 225,
		"New York": 225,
		"Florida": 225,
		"Ohio": 225,
		"Texas": 225,
	}
	for k, v := range statePopulation {
		fmt.Println(k, v)
	}
	fmt.Println("=====================")

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
