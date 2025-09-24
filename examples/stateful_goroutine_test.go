package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func TestStatefulGoroutine(t *testing.T) {

	t.Run("stateful goroutine", func(t *testing.T) {
		var readOps uint64
		var writeOps uint64

		reads := make(chan readOp)
		writes := make(chan writeOp)
		flag := make(chan bool, 1)

		go func() {
			var state = make(map[int]int)
			for {
				select {
				case read := <-reads:
					val := state[read.key]
					read.resp <- val
				case write := <-writes:
					state[write.key] = write.value
					write.resp <- true
				case <-flag:
					fmt.Println(state)
				}
			}
		}()

		var wg sync.WaitGroup
		wg.Add(100)
		for range 100 {
			go func() {
				for range 100 {
					read := readOp{
						key:  rand.Intn(5),
						resp: make(chan int),
					}
					reads <- read
					<-read.resp
					atomic.AddUint64(&readOps, 1)
					time.Sleep(time.Millisecond)
				}
				wg.Done()
			}()
		}
		wg.Add(10)
		for range 10 {
			go func() {
				for range 10 {
					write := writeOp{
						key:   rand.Intn(5),
						value: rand.Intn(100),
						resp:  make(chan bool),
					}
					writes <- write
					<-write.resp
					atomic.AddUint64(&writeOps, 1)
					time.Sleep(time.Millisecond)
				}
				wg.Done()
			}()
		}

		time.Sleep(time.Second)
		wg.Wait()
		flag <- true
		readOpsFinal := atomic.LoadUint64(&readOps)
		fmt.Println("readOps", readOpsFinal)
		writeOpsFinal := atomic.LoadUint64(&writeOps)
		fmt.Println("writeOps", writeOpsFinal)
	})
}
