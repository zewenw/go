package examples

import (
	"fmt"
	"testing"
)

type Application struct {
	name string
}

func send(app Application) {
	fmt.Printf("inner: %p\n", &app)
	app.name = "world"
	fmt.Println(app)
}

func TestPassing(t *testing.T) {
	t.Run("reference", func(t *testing.T) {
		app := &Application{
			name: "hello",
		}
		fmt.Println("outter:", &app)
		send(*app)
		fmt.Println(*app)
	})
}

func TestArray(t *testing.T) {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 4: 400}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [...][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
