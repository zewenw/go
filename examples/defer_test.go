package examples

import (
	"fmt"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {

	t.Run("defer demo", func(t *testing.T) {
		f := CreatFile("defer.txt")
		defer RemoveFile(f)
		defer CloseFile(f)
		WriteFile(f)
	})
}

func CreatFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func WriteFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func CloseFile(f *os.File) {
	fmt.Println("Closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func RemoveFile(f *os.File) {
	fmt.Println("Delete File")
	err := os.Remove(f.Name())
	if err != nil {
		panic(err)
	}
}
