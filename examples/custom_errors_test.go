package examples

import (
	"errors"
	"fmt"
	"testing"
)

type argError struct {
	arg     int
	message string
}

func (ae *argError) Error() string {
	return fmt.Sprintf("%d - %s", ae.arg, ae.message)
}

func fe(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func TestCustomError(t *testing.T) {

	t.Run("custom error", func(t *testing.T) {
		_, err := fe(42)
		var ae *argError
		if errors.As(err, &ae) {
			fmt.Println(ae.arg)
			fmt.Println(ae.Error())
		} else {
			fmt.Println("error doesn't match argError")
		}
	})
}
