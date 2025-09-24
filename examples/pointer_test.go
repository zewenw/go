package examples

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {

	t.Run("pointer test", func(t *testing.T) {
		val := 1
		fmt.Println("initial address:", &val)
		zeroval(val)
		fmt.Println("val:", val)
		ptrZeroVal(&val)
		fmt.Println("val:", val)

		fmt.Println("Pointer:", &val)
	})

}

func zeroval(ival int) {
	ival = 0
	fmt.Println("inner address:", &ival)
}

func ptrZeroVal(ival *int) {
	*ival = 0
}
