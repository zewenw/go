package examples

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestRandomNum(t *testing.T) {

	t.Run("number parsing", func(t *testing.T) {
		f, _ := strconv.ParseFloat("1.234", 64)
		fmt.Println(f)

		i, _ := strconv.ParseInt("123", 0, 64)
		fmt.Println(i)

		d, _ := strconv.ParseInt("0x1c8", 0, 64)
		fmt.Println(d)

		u, _ := strconv.ParseUint("789", 0, 64)
		fmt.Println(u)

		k, _ := strconv.Atoi("135")
		fmt.Println(k)

		_, e := strconv.Atoi("wat")
		fmt.Println(e)

	})

	t.Run("Random Number", func(t *testing.T) {
		fmt.Println(rand.Intn(100), ",")
		fmt.Println(rand.Intn(100))

		fmt.Println()

		fmt.Println(rand.Float64())
		fmt.Println((rand.Float64()*5)+5, ",")

		s1 := rand.NewSource(7)
		fmt.Println(s1.Int63(), ",")

	})
}
