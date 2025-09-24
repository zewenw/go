package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiting(t *testing.T) {

	t.Run("Rate lImiting", func(t *testing.T) {
		requests := make(chan int, 5)
		for i := range 5 {
			requests <- i
		}
		close(requests)

		limiter := time.Tick(200 * time.Millisecond)
		for req := range requests {
			t := <-limiter
			fmt.Println("Request:", req, t)
		}
	})

	t.Run("Burst Limiter", func(t *testing.T) {
		burstLimiter := make(chan time.Time, 3)
		for range 3 {
			burstLimiter <- time.Now()
		}

		go func() {
			for t := range time.Tick(200 * time.Millisecond) {
				burstLimiter <- t
			}
		}()

		burstRequests := make(chan int, 5)
		for i := range 5 {
			burstRequests <- i
		}
		close(burstRequests)
		for req := range burstRequests {
			t := <-burstLimiter
			fmt.Println("request", req, t)
		}

	})
}
