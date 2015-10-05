package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func cleanName(name string) string {
	// clean - & _ from filename
	name = strings.Replace(name, "-", "", -1)

	// capitalize the first letter
	a := []rune(name)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(wd, "icons", "txt")
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	out, err := os.Create(filepath.Join(wd, "icons", "icons.go"))
	if err != nil {
		panic(err)
	}
	out.Write([]byte("// AUTOGENERATED FILE; DO NOT EDIT DIRECTLY\n// See icons/generate/generate.go for more info\npackage icons\n\nconst (\n"))
	for _, f := range fs {
		// get all the text files
		if strings.HasSuffix(f.Name(), ".txt") {
			name := cleanName(f.Name())
			out.Write([]byte(strings.TrimSuffix(name, ".txt") + " = `"))
			f, err := os.Open(filepath.Join(path, f.Name()))
			if err != nil {
				panic(err)
			}
			io.Copy(out, f)
			f.Close()
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
