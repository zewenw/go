package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {

	t.Run("channel demo", func(t *testing.T) {
		message := make(chan string)

		go func() {
			message <- "ping"
		}()

		msg := <-message
		fmt.Println("message got from channel:", msg)
	})

	t.Run("buffered channel", func(t *testing.T) {
		messages := make(chan string, 1)

		messages <- "buffered"
		messages <- "channel"
		fmt.Println(<-messages)
		messages <- "third"
		fmt.Println("the third value was sent")

		fmt.Println(<-messages)
	})

	t.Run("channel synchronization", func(t *testing.T) {
		done := make(chan bool, 1)
		go worker(done)
		<-done
	})

	t.Run("channel direction", func(t *testing.T) {
		pings := make(chan string, 1)
		pongs := make(chan string, 1)
		ping(pings, "msg for ping channel")
		pong(pings, pongs)
		fmt.Println(<-pongs)
	})
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(ping <-chan string, pongs chan<- string) {
	msg := <-ping
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Println("working ...")
	time.Sleep(time.Second)
	fmt.Println("working is done")
	done <- true
}
