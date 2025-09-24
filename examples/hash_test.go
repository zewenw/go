package examples

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	t.Run("Hash demo", func(t *testing.T) {
		s := "sha256 this string"

		h := sha256.New()
		h.Write([]byte(s))
		bs := h.Sum(nil)
		fmt.Println(s)
		fmt.Printf("%x\n", bs)
	})
}
