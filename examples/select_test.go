package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestSelct(t *testing.T) {

	t.Run("select demo", func(t *testing.T) {
		c1 := make(chan string)
		c2 := make(chan string)
		go func() {
			time.Sleep(1 * time.Second)
			c1 <- "one"
		}()

		go func() {
			time.Sleep(1 * time.Second)
			c2 <- "two"
		}()

		for i := 0; i < 2; i++ {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	})
}
