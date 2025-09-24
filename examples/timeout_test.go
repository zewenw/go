package examples

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {

	t.Run("timeout demo", func(t *testing.T) {
		c1 := make(chan string)
		go func() {
			time.Sleep(2 * time.Second)
			c1 <- "Result from channel one"
		}()

		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(1 * time.Second):
			fmt.Println("timeout for channel one")
		}

		c2 := make(chan string)
		go func() {
			time.Sleep(2 * time.Second)
			c2 <- "Result from channel two"
		}()

		select {
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
		}
	})

	t.Run("non blocking channel", func(t *testing.T) {
		messages := make(chan string)
		signals := make(chan string)

		select {
		case msg := <-messages:
			fmt.Println("message received: ", msg)
		default:
			fmt.Println("no message received")
		}

		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("send message", msg)
		default:
			fmt.Println("no message sent")
		}

		select {
		case msg := <-messages:
			fmt.Println("message received", msg)
		case sig := <-signals:
			fmt.Println("received signal", sig)
		default:
			fmt.Println("no activity")
		}
	})

	t.Run("Closing channel Demo", func(t *testing.T) {
		jobs := make(chan string, 5)
		done := make(chan bool)

		go func() {
			for {
				job, more := <-jobs
				if more {
					fmt.Println("received job", job)
				} else {
					fmt.Println("receive all jobs")
					done <- true
					return
				}
			}
		}()

		for i := range 5 {
			jobs <- strconv.Itoa(i)
			fmt.Println("send job:", i)
		}
		close(jobs)
		fmt.Println("sent all jobs")
		<-done

		_, ok := <-jobs
		fmt.Println("received more job:", ok)
	})

	t.Run("iterate over a channel", func(t *testing.T) {
		queue := make(chan string, 2)
		queue <- "one"
		queue <- "two"

		close(queue)
		for elem := range queue {
			fmt.Println(elem)
		}
	})

	t.Run("ticker demo", func(t *testing.T) {
		ticker := time.NewTicker(500 * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case t := <-ticker.C:
					fmt.Println("tick at", t)
				case <-done:
					return
				}
			}
		}()

		time.Sleep(1600 * time.Millisecond)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker Stopped")
	})
}
