package main

import (
	"fmt"
)

func main(){
	statePopulation := make(map[string]int)
	// statePopulation := map[string]int{
	statePopulation = map[string]int{
		"California": 225,
		"New York": 225,
		"Florida": 225,
		"Ohio": 225,
		"Texas": 225,
	}
	statePopulation["Georgia"] = 5352
	fmt.Println(statePopulation["Georgia"])
	delete(statePopulation, "Georgia")
	fmt.Println(statePopulation["Georgia"])

	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Ohio"])

	//misspelling which the corresponding value shouldn't be 0
	pop, ok := statePopulation["Ohio"]
	// pop, ok := statePopulation["Oho"]

	// throw away the first result object
	// _, ok := statePopulation["Ohio"]
	
	fmt.Println(pop, ok)

	fmt.Println(len(statePopulation))
}