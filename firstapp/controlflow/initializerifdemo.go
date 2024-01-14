package main

import (
	"fmt"
)

func main(){
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 225,
		"New York": 225,
		"Florida": 225,
		"Ohio": 225,
		"Texas": 225,
	}
	if pop, ok := statePopulation["Florida"]; ok {
		fmt.Println(pop)
	}
}