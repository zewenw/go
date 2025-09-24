package examples

import (
	"fmt"
	"runtime"
	"testing"
)

func captureStackTrace() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

func myPanic() error {
	// Simulate a panic
	panic("a problem")
}

func safeFunction() (err error) {
	defer func() {
		if r := recover(); r != nil {
			stackTrace := captureStackTrace()
			if e, ok := r.(error); ok {
				err = fmt.Errorf("Recovered from panic: %w\nStack trace:\n%s", e, stackTrace)
			} else {
				err = fmt.Errorf("Recovered from panic: %v\nStack trace:\n%s", r, stackTrace)
			}
		}
	}()

	// Call a function that panics
	myPanic()
	return nil
}

func TestRecover(t *testing.T) {
	t.Run("demo", func(t *testing.T) {
		err := safeFunction()
		if err != nil {
			fmt.Println("Error captured with stack trace:")
			fmt.Println(err)
		}
		fmt.Println("Recovered, Error:\n", err)
		fmt.Println("After myPanic()")
	})

}
