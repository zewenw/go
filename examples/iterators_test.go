package examples

import (
	"fmt"
	"iter"
	"testing"
)

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func getFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func TestIterator(t *testing.T) {

	t.Run("Iteration Deom", func(t *testing.T) {
		lst := List[int]{}
		lst.Push(10)
		lst.Push(12)
		lst.Push(13)
		//for e := range lst.All() {
		//	fmt.Println(e)
		//}
		yieldFunc := lst.All()

		for n := range getFib() {
			if n >= 10 {
				break
			}
			fmt.Println(n)
		}

		var container []int
		yieldFunc(func(val int) bool {
			fmt.Println(val)
			container = append(container, val)
			if val == 12 {
				return false
			}
			return true
		})
		fmt.Println(container)
	})
}
