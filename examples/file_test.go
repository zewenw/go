package examples

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestFile(t *testing.T) {

	t.Run("Writing File", func(t *testing.T) {
		d1 := []byte("hello\ngo\n")
		err := os.WriteFile("./dat1", d1, 0644)
		check(err)

		f, err := os.Create(".dat2")

		defer f.Close()

		d2 := []byte{115, 111, 109, 101, 10}
		n2, err := f.Write(d2)
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		n3, err := f.WriteString("writes\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n3)

		f.Sync()

		w := bufio.NewWriter(f)
		n4, err := w.WriteString("buffered\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n4)

		w.Flush()

	})

	t.Run("Reading File", func(t *testing.T) {
		fileName := "../examples/array_test.go"
		data, err := os.ReadFile(fileName)
		check(err)
		fmt.Println(string(data))

		f, err := os.Open(fileName)
		defer f.Close()
		check(err)

		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

		o2, err := f.Seek(6, io.SeekStart)
		check(err)
		b2 := make([]byte, 2)
		n2, err := f.Read(b2)
		check(err)
		fmt.Printf("%d bytes @ %d: ", n2, o2)
		fmt.Printf("%v\n", string(b2[:n2]))

		_, err = f.Seek(2, io.SeekCurrent)
		check(err)

		_, err = f.Seek(-4, io.SeekEnd)
		check(err)

		o3, err := f.Seek(6, io.SeekStart)
		check(err)
		b3 := make([]byte, 2)
		n3, err := io.ReadAtLeast(f, b3, 2)
		check(err)
		fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

		_, err = f.Seek(0, io.SeekStart)
		check(err)

		r4 := bufio.NewReader(f)
		b4, err := r4.Peek(5)
		check(err)
		fmt.Printf("5 bytes: %s\n", string(b4))
	})
}
