package examples

import (
	"fmt"
	"sync"
	"testing"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func TestMutex(t *testing.T) {

	t.Run("Mutext Demo", func(t *testing.T) {
		counter := Container{
			counters: map[string]int{"a": 0, "b": 0},
		}

		var wg sync.WaitGroup
		doIncrement := func(name string, n int) {
			for range n {
				counter.inc(name)
			}
			wg.Done()
		}

		wg.Add(3)
		go doIncrement("a", 1000)
		go doIncrement("a", 1000)
		go doIncrement("a", 1000)
		wg.Wait()
		fmt.Println(counter.counters)
	})
}
