package main

import (
	"fmt"
)

func main(){
	
	var divide func (float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cann't divide by zero")
		}
		return a / b, nil
	}
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(d)


	fmt.Println("========================")
	f := func () {
		fmt.Println("Hello GOLANG")
	}
	f()

	fmt.Println("========================")
	for i := 0; i < 5; i++ {
		func (i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("========================")

	//anonymous function
	func(){
		fmt.Println("hello go")
	}()
}