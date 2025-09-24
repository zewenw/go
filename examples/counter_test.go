package examples

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("Counter Demo", func(t *testing.T) {
		var ops uint

		var wg sync.WaitGroup

		for range 50 {
			wg.Add(1)
			go func() {
				for range 1000 {
					ops++
				}
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Println("ops:", ops)
	})
}
