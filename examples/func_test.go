package examples

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	t.Run("func test", func(t *testing.T) {
		res := plus(1, 2)
		fmt.Println("plus:", res)

		res1 := plusPlus(1, 2, 3)
		fmt.Println("plusPlus", res1)
	})
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func TestMultipleReturns(t *testing.T) {
	t.Run("multiple returns", func(t *testing.T) {
		a, b := vals()
		fmt.Println("val:", a)
		fmt.Println("val:", b)

		_, c := vals()
		fmt.Println("val:", c)
	})
}

func vals() (int, int) {
	return 3, 7
}

func TestVariadicFunc(t *testing.T) {
	t.Run("VariadicFunc", func(t *testing.T) {
		sum(1, 2)
		sum(1, 2, 3)

		nums := []int{1, 2, 3, 4}
		sum(nums...)
	})
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, val := range nums {
		total += val
	}
	fmt.Println(total)
}

func TestClosure(t *testing.T) {
	t.Run("Closure function", func(t *testing.T) {
		nextInt := intSeq()

		fmt.Println(nextInt())
		fmt.Println(nextInt())
		fmt.Println(nextInt())

		newInt := intSeq()
		fmt.Println(newInt())
	})
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestRecursion(t *testing.T) {
	t.Run("Testing Recursion", func(t *testing.T) {
		fmt.Println(fact(5))
	})

	t.Run("Anonymous Testing", func(t *testing.T) {
		var fib func(n int) int
		fib = func(n int) int {
			if n < 2 {
				return n
			}

			return fib(n-1) + fib(n-2)
		}
		fmt.Println(fib(7))
	})
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
