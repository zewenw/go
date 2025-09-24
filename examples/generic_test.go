package examples

import (
	"fmt"
	"testing"
)

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func TestGeneric(t *testing.T) {

	t.Run("generics demo", func(t *testing.T) {
		var s = []string{"foo", "bar", "zoo"}

		fmt.Println("the index of foo:", SliceIndex(s, "foo"))

		lst := List[int]{}
		lst.Push(10)
		lst.Push(13)
		lst.Push(23)
		fmt.Println("list:", lst.AllElements())
	})
}
