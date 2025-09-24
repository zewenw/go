package examples

import (
	"fmt"
	"testing"
)

func TestRangeFunc(t *testing.T) {
	t.Run("range slice", func(t *testing.T) {
		nums := []int{1, 2, 3}
		sum := 0
		for _, num := range nums {
			sum += num
		}
		fmt.Println("sum:", sum)

		for i, num := range nums {
			if num == 3 {
				fmt.Println("index:", i)
			}
		}
	})

	t.Run("range map", func(t *testing.T) {
		kvs := map[string]string{"a": "apple", "b": "banana"}
		for k, v := range kvs {
			fmt.Printf("%s -> %s\n", k, v)
		}

		for k := range kvs {
			fmt.Println("key:", k)
		}
	})

	t.Run("Range Rune", func(t *testing.T) {
		for i, c := range "go" {
			fmt.Println(i, c)
		}
	})

}
