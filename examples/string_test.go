package examples

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode/utf8"
)

type point struct {
	x, y int
}

func TestStrings(t *testing.T) {

	t.Run("float and int", func(t *testing.T) {
		var f float64 = 123
		fmt.Println(f)
	})

	t.Run("String formatting", func(t *testing.T) {
		var p = fmt.Printf
		po := point{1, 2}
		p("struct1: %v\n", po)
		p("struct2: %+v\n", po)
		p("struct3: %#v\n", po)
		p("type: %T\n", po)
		p("bool: %t\n", true)
		p("int: %d\n", 123)
		p("binary: %b\n", 14)
		p("char: %c\n", 33)
		p("hex: %x\n", 456)
		p("float1: %f\n", 78.9)
		p("float2: %e\n", 123400000.0)
		p("float2: %E\n", 123400000.0)

		p("str1: %s\n", "\"string\"")
		p("str2: %q\n", "\"string\"")
		p("str3: %x\n", "hex this")

		p("pointer: %p\n", &p)
		p("width1: |%6d|%6d|\n", 12, 345)
		p("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
		p("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
		p("width4: |%6s|%6s|\n", "foo", "b")
		p("width5: |%-6s|%-6s|\n", "foo", "b")

		s := fmt.Sprintf("sprintf: a %s", "string")
		fmt.Println(s)

		fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	})

	t.Run("string functions", func(t *testing.T) {
		var p = fmt.Println
		p("Contains:", strings.Contains("test", "es"))
		p("Count:", strings.Count("test", "t"))
		p("HasPrefix:", strings.HasPrefix("test", "te"))
		p("HasSuffix:", strings.HasSuffix("test", "st"))
		p("Index:", strings.Index("test", "s"))
		p("Join:", strings.Join([]string{"a", "b"}, "c"))
		p("Repeat:", strings.Repeat("s", 5))
		p("Replace:", strings.Replace("foo", "o", "x", -1))
		p("Replace:", strings.Replace("foo", "o", "x", 1))
		p("Split:", strings.Split("a-b-c-d-e", "-"))
		p("Splitn:", strings.SplitN("a-b-c-d-e", "-", 2))
		p("ToLower:", strings.ToLower("ABC"))
		p("ToUpper:", strings.ToUpper("abc"))
	})

	t.Run("String test", func(t *testing.T) {
		const s = "สวัสดี"
		fmt.Println("len:", len(s))

		for i := 0; i < len(s); i++ {
			fmt.Printf("%x ", s[i])
		}
		fmt.Println()

		fmt.Println("Rune Count:", utf8.RuneCountInString(s))

		for idx, runeValue := range s {
			fmt.Printf("%#U starts at %d\n", runeValue, idx)
		}
		fmt.Println("\nUsing DecodeRuneInString")

		for i, w := 0, 0; i < len(s); i += w {
			runeValue, width := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%#U starts at %d\n", runeValue, i)
			w = width
			examineRune(runeValue)
		}
	})
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
