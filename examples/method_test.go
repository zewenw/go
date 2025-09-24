package examples

import (
	"fmt"
	"testing"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	fmt.Printf("Area address: %p\n", r)
	return r.width * r.height
}

func (r *rect) perim() int {
	fmt.Printf("Perim address: %p\n", &r)
	return 2*r.width + 2*r.height
}

func TestMethod(t *testing.T) {

	t.Run("Method Test", func(t *testing.T) {
		r := rect{width: 10, height: 10}
		fmt.Printf("address: %p\n", &r)
		fmt.Println("area: ", r.area())
		fmt.Println("Perim: ", r.perim())

		rp := &r
		fmt.Println("area: ", rp.area())
		fmt.Println("Perim: ", rp.perim())
	})
}
