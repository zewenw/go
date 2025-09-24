package examples

import (
	"cmp"
	"fmt"
	"slices"
	"testing"
)

func TestSorting(t *testing.T) {

	t.Run("sorting", func(t *testing.T) {
		strs := []string{"c", "a", "b"}
		slices.Sort(strs)
		fmt.Println("Strings:", strs)

		ints := []int{3, 2, 4}
		slices.Sort(ints)
		fmt.Println("Ints:", ints)

		s := slices.IsSorted(ints)
		fmt.Println("Sorted:", s)
	})

	t.Run("Sorting by function", func(t *testing.T) {
		fruits := []string{"peach", "banana", "kiwi"}
		lenCmp := func(a, b string) int {
			return cmp.Compare(len(a), len(b))
		}
		slices.SortFunc(fruits, lenCmp)
		fmt.Println(fruits)

		type Person struct {
			name string
			age  int
		}
		people := []*Person{
			{name: "Jax", age: 37},
			{name: "TJ", age: 25},
			{name: "Alex", age: 72},
		}
		slices.SortFunc(people, func(a, b *Person) int {
			return a.age - b.age
		})
		fmt.Println(people)
	})
}
