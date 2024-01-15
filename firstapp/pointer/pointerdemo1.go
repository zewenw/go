package main

import (
	"fmt"
)

func main (){
	var ms *myStruct
	// initial value for pointer is nil
	fmt.Println(ms)
	ms = new(myStruct)
	ms.foo = 42
	// ms = &myStruct{foo: 42}
	// fmt.Println(ms)
	fmt.Println((*ms).foo)
}

type myStruct struct{
	foo int
}