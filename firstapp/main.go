package main

import "fmt"

var j float32 = 6.0

var (
	actorName string = "Elon"
	companion string = "sarah"
	doctorNumber int = 3
	season = 11
)


func main(){
	i := 42
	fmt.Println(i)
	fmt.Printf("%v, %T", i , i)
	fmt.Printf("%v, %T", j , j)
}
