package examples

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func workerint(id int) {
	fmt.Printf("Worder %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Woker %d done\n", id)
}

func TestWaitGroup(t *testing.T) {
	t.Run("wait group demo", func(t *testing.T) {
		var wg sync.WaitGroup

		for i := 1; i <= 5; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				workerint(i)
			}()
		}
		wg.Wait()
		fmt.Println("All works are finished")
	})
}
