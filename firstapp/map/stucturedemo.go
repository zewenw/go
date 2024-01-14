package main

import (
	"fmt"
)

type Doctor struct {
	//outside only know this struct but the field name a private unless we change the leading letter to uppercase
	// Number     int
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Joe Pertee",
		companions: []string{
			"Liz",
			"Sarah",
		},
	}
	//position syntax need author give value in correct order
	aDoctor1 := Doctor{
		3,
		"Joe Pertee",
		[]string{
			"Liz",
			"Sarah",
		},
	}
	fmt.Println(aDoctor1)
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
}
