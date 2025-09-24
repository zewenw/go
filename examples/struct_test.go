package examples

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{
		name: name,
	}
	p.age = 42
	return &p
}

func TestStruct(t *testing.T) {
	t.Run("testing struct", func(t *testing.T) {
		fmt.Println(person{"Bob", 20})

		fmt.Println(person{
			name: "Alice",
			age:  30,
		})

		fmt.Println(person{name: "Fred"})

		fmt.Println(&person{name: "Ann", age: 40})

		fmt.Println(newPerson("John"))

		s := person{name: "Sean", age: 50}
		fmt.Println(s.name)

		sp := &s
		fmt.Println(sp.age)

		sp.age = 51
		fmt.Println(s.age)

		dog := struct {
			name   string
			isGood bool
		}{
			"Rex",
			true,
		}
		fmt.Println(dog)
	})
}
