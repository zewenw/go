package examples

import (
	"errors"
	"fmt"
	"testing"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("my Error")
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		myError := &myError{
			"hello",
		}
		return fmt.Errorf("making tea: %w", myError)
	}
	return nil
}

func TestError(t *testing.T) {

	t.Run("error check", func(t *testing.T) {
		var myE *myError
		for i := range 5 {
			if err := makeTea(i); err != nil {
				if errors.Is(err, ErrOutOfTea) {
					fmt.Println("We should buy new tea!")
				} else if errors.As(err, &myE) {
					fmt.Println(myE)
					fmt.Println("Now is dark.")
				} else {
					fmt.Printf("Unknown Error, %s\n", err)
				}
				continue
			}
			fmt.Println("Tea is ready")
		}
	})

	t.Run("Error Test", func(t *testing.T) {
		for _, i := range []int{7, 42} {
			if r, e := f(i); e != nil {
				fmt.Println("f failed:", e)
			} else {
				fmt.Println("f worked:", r)
			}
		}
	})
}
