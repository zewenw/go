package examples

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	t.Run("signla demo", func(t *testing.T) {
		sigs := make(chan os.Signal, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

		done := make(chan bool, 1)

		go func() {

			sig := sigs
			fmt.Println()
			fmt.Println("Signal:", sig)
			time.Sleep(5 * time.Second)
			done <- true
		}()

		fmt.Println("awaiting signal")
		<-done
		fmt.Println("exiting")

	})

	t.Run("map without key-value pair", func(t *testing.T) {
		demo := map[string]string{
			"key":  "value",
			"key2": " ",
		}
		var customerId string = "customerID"
		if demo["key2"] == customerId {
			fmt.Println("found it")
		} else {
			fmt.Println("Not exist")
		}

	})
}
