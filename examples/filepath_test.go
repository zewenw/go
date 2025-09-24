package examples

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFilePath(t *testing.T) {

	var check = func(e error) {
		if e != nil {
			panic(e)
		}
	}

	t.Run("Temp file", func(t *testing.T) {
		f, err := os.CreateTemp("", "sample")
		check(err)

		fmt.Println("Temp file name:", f.Name())

		defer os.Remove(f.Name())

		_, err = f.Write([]byte{1, 2, 3, 4})
		check(err)

		dname, err := os.MkdirTemp("", "sampledir")
		check(err)
		fmt.Println("Temp dir name:", dname)

		defer os.RemoveAll(dname)

		fname := filepath.Join(dname, "file1")
		err = os.WriteFile(fname, []byte{1, 2}, 0666)
		check(err)

	})

	t.Run("Dir", func(t *testing.T) {
		err := os.Mkdir("subdir", 0755)
		check(err)
		defer os.RemoveAll("subdir")

		createEmptyFile := func(name string) {
			d := []byte("")
			check(os.WriteFile(name, d, 0644))
		}
		createEmptyFile("subdir/file1")

		err = os.MkdirAll("subdir/parent/child", 0755)
		check(err)

		createEmptyFile("subdir/parent/file2")
		createEmptyFile("subdir/parent/file3")
		createEmptyFile("subdir/parent/child/file4")

		c, err := os.ReadDir("subdir/parent")
		check(err)
		fmt.Println("Listing subdir/parent")
		for _, entry := range c {
			fmt.Println(" ", entry.Name(), entry.IsDir())
		}

		err = os.Chdir("subdir/parent/child")
		check(err)

		c, err = os.ReadDir(".")
		check(err)
		fmt.Println("Listing subdir/parent/child")
		for _, entry := range c {
			fmt.Println(" ", entry.Name(), entry.IsDir())
		}

		err = os.Chdir("../../..")
		check(err)

		walkDir := func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fmt.Println("	", path, d.IsDir())
			return nil
		}

		fmt.Println("Visiting subdir")
		err = filepath.WalkDir("subdir", walkDir)

	})

	t.Run("File Path", func(t *testing.T) {
		p := filepath.Join("dir1", "dir2", "filename")
		fmt.Println("p:", p)

		fmt.Println(filepath.Join("dir1//", "filename"))
		fmt.Println(filepath.Join("dir1/.../dir1", "filename"))

		fmt.Println("Dir(p):", filepath.Dir(p))
		fmt.Println("Base(p):", filepath.Base(p))

		filename := "config.json"

		ext := filepath.Ext(filename)
		fmt.Println(ext)

		fmt.Println(strings.TrimSuffix(filename, ext))

		rel, err := filepath.Rel("a/b", "a/b/t/file")
		if err != nil {
			panic(err)
		}
		fmt.Println(rel)

		rel, err = filepath.Rel("a/b", "a/c/t/file")
		if err != nil {
			panic(err)
		}
		fmt.Println(rel)
	})
}
