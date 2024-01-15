package main

import(
	"fmt"
	"log"
)

func main(){
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker(){
	fmt.Println("about to panic")
	defer func(){
		if err := recover(); err != nil {
			// following code will pass exception to calling function
			// log.Panicln("Error:", err)
			// panic(err)
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}