package examples

import (
	"fmt"
	"testing"
	"time"
)

func f1(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}

func TestGoroutine(t *testing.T) {

	t.Run("Easy Goroutine", func(t *testing.T) {
		f1("direct")

		go f1("goroutine")

		go func(msg string) {
			fmt.Println(msg)
		}("anonymous")

		time.Sleep(time.Second)

		fmt.Println("Done")
	})

	t.Run("Return in routine", func(t *testing.T) {
		go func() {
			fmt.Println("start return")
			fmt.Println("after return")
		}()
	})
}
