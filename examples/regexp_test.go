package examples

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {

	t.Run("Regexp demo", func(t *testing.T) {
		match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
		fmt.Println(match)

		r, _ := regexp.Compile("p([a-z]+)ch")
		fmt.Println(r.MatchString("peach"))

		fmt.Println(r.FindString("peach punch"))

		fmt.Println("idx", r.FindStringIndex("peach punch"))

		fmt.Println("sub match", r.FindStringSubmatch("peach punch"))

		fmt.Println("all", r.FindAllString("peach punch pinch", -1))

		fmt.Println("all:", r.FindAllStringSubmatchIndex(
			"peach punch pinch", -1))

		fmt.Println(r.FindAllString("peach punch pinch", 2))

		fmt.Println(r.Match([]byte("peach")))

		r = regexp.MustCompile("p([a-z]+)ch")
		fmt.Println("regexp:", r)

		fmt.Println(r.ReplaceAllString("a peach", "<Fruit>"))

		in := []byte("a peach")
		out := r.ReplaceAllFunc(in, bytes.ToUpper)
		fmt.Println(string(out))
	})
}
