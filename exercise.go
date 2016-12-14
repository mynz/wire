package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	// "strconv"
)

type HashDigit [64]byte

func (hash *HashDigit) String() string {
	str := ""
	for _, d := range hash {
		str += fmt.Sprintf("%02x", d)
	}
	return str
}

////

type FileSpec struct {
	Path       string
	FileInfo   os.FileInfo
	HashString string
}

func NewFileSpec(path string, fi os.FileInfo) *FileSpec {
	return &FileSpec{path, fi, ""}
}

////

func main() {
	fmt.Println("Hello")

	fileSpecs := make([]FileSpec, 0)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
		}
		// check if it is a relgular file.
		if info.Mode().IsRegular() {
			fileSpecs = append(fileSpecs, *NewFileSpec(path, info))
		}
		return nil
	})

	for i, fspec := range fileSpecs {

		path := fspec.Path

		fmt.Println("file", i, path)

		if bs, err := ioutil.ReadFile(path); err != nil {
			panic(err)
		} else {
			fmt.Println(len(bs))

			hex := HashDigit(sha512.Sum512(bs))
			fmt.Println("hex", hex)

			hstr := hex.String()
			fmt.Println(hstr)

		}
	}
}
