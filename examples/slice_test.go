package examples

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("app", s)
	s = append(s, "e", "f")
	fmt.Println("app:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cop:", c)
	fmt.Println("val:", c[0] == s[0], "add:", &c[0] == &s[0])

	l := s[2:3]
	fmt.Println("len sl2:", len(l))
	fmt.Println("cap sl2:", cap(l))

	l = s[2:]
	fmt.Println("sl3:", l)

	s1 := []string{"g", "h", "i"}
	fmt.Println("dcl:", s1)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
