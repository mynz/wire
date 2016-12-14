package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	// "strings"
)

func main() {
	fmt.Println("Hello")

	type FileSpec struct {
		Path     string
		FileInfo os.FileInfo
	}

	fileSpecs := make([]FileSpec, 0)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
		}
		// check if it is a relgular file.
		if info.Mode().IsRegular() {
			fileSpecs = append(fileSpecs, FileSpec{path, info})
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
		}

	}

}
