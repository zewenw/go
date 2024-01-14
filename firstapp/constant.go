package main

import "fmt"

// inner variable shadow package variable and can change the type of package variable
const a int16 = 27

const (
	// iota is a counter
	j = iota
	k = iota
)

func main(){
	// leading by a uppercase letter will export to outside
	// const MyCont
	const myConst int = 42

	/*
	if we changes the value of a const then it will trigger a error
	*/
	// myConst = 25
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("=============================\n")

	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	const e rune = 'b'
	const f = 'b'

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	fmt.Printf("=============================\n")
	fmt.Printf("%v\n", j)
	fmt.Printf("%v\n", k)
}