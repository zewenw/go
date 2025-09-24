package examples

import (
	"fmt"
	"testing"
	"time"
)

func TestEpoch(t *testing.T) {

	t.Run("Epoch demo", func(t *testing.T) {
		now := time.Now()
		fmt.Println(now)

		fmt.Println(now.Unix())
		fmt.Println(now.UnixMilli())
		fmt.Println(now.UnixNano())

		fmt.Println(time.Unix(now.Unix(), 0))
		fmt.Println(time.Unix(0, now.UnixNano()))
	})
}
