package main

import(
	"fmt"
)

func main(){
	aDoctor := struct{name string}{name: "john Pertwee"}
	//copy of aDoctor
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
	
	//passing by reference
	bDoctor := &aDoctor
	bDoctor.name = "Jack"
	fmt.Println(aDoctor)
	fmt.Println(bDoctor)
}