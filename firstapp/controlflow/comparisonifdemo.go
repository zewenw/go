package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 130

	// if guess < 1 || returnTrue() || guess > 100 {
	if guess < 1 {
		fmt.Println("The guess must be greater than 1")
	} else if guess > 100 {
		fmt.Println("The guess must be less than 100")
	} else {
		if guess < number {
			fmt.Println("too low")
		}
		if guess > number {
			fmt.Println("too high")
		}
		if guess == number {
			fmt.Println("you are right")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	fmt.Println(!true)
}

func returnTrue() bool {
	fmt.Println("return true")
	return true
}
