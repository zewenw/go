package examples

import (
	"fmt"
	"math"
	"testing"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}

func (r rect1) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius:", c.radius)
	}
	fmt.Println("this is not a circle")
}

func TestInterface(t *testing.T) {

	t.Run("interface test", func(t *testing.T) {
		r := rect1{width: 3, height: 5}
		c := circle{radius: 2}
		r1 := geometry(r)
		fmt.Println(r1)
		if k, ok := r1.(rect1); ok {
			fmt.Println("this is a rec type", k)
		}
		measure(r)
		measure(c)

		detectCircle(r)
		detectCircle(c)

	})
}
