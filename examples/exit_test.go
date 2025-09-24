package examples

import (
	"fmt"
	"os"
	"testing"
)

func TestExit(t *testing.T) {

	t.Run("Exit test", func(t *testing.T) {
		defer fmt.Println("!")

		os.Exit(3)
	})
}
