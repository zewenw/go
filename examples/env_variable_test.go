package examples

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnvVariables(t *testing.T) {

	t.Run("env variables", func(t *testing.T) {
		os.Setenv("FOO", "1")
		fmt.Println("FOO", os.Getenv("FOO"))
		fmt.Println("BAR", os.Getenv("BAR"))

		fmt.Println()
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Println(pair[0])
		}
	})
}
