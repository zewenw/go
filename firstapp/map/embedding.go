package main

import(
	"fmt"
)

type Animal struct{
	Name string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

func main (){
	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 45
	// b.CanFly = false
	b := Bird {
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 45,
		CanFly: false,
	}
	fmt.Println(b)
}

