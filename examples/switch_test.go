package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitch(t *testing.T) {
	t.Run("switch demo", func(t *testing.T) {
		switch time.Sunday {
		case time.Saturday, time.Sunday:
			fmt.Println("Weekend")
		default:
			fmt.Println("Working day")
		}
	})

	t.Run("type assertion", func(t *testing.T) {
		whatAmI := func(i interface{}) {
			switch t := i.(type) {
			case bool:
				fmt.Println("I am a boolean")
			case int:
				fmt.Println("I am a integer")
			default:
				fmt.Printf("Don't know type %T\n", t)
			}
		}
		whatAmI(1)
		whatAmI(true)
		whatAmI(0.01)
	})
}
