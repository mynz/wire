package wire

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
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
	Path  string    `json:"path"`
	Size  int64     `json:"size"`
	MTime time.Time `json:"mtime"`
	Hash  string    `json:"hash"`
}

func (fs FileSpec) String() string {
	return fmt.Sprintf("path: %s, size: %d, Hash: %s...", fs.Path, fs.Size, fs.Hash[0:8])
}

////

func CollectFileSpecs(rootDir string) []FileSpec {

	type Info struct {
		Path     string
		FileInfo os.FileInfo
	}

	infoList := make([]Info, 0)

	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
		}
		// check if it is a relgular file.
		if info.Mode().IsRegular() {
			fs := Info{Path: path, FileInfo: info}
			infoList = append(infoList, fs)
		}
		return nil
	})

	specs := make([]FileSpec, 0, len(infoList))
	for _, info := range infoList {
		path := info.Path
		size := info.FileInfo.Size()
		mtime := info.FileInfo.ModTime()
		if bs, err := ioutil.ReadFile(path); err != nil {
			panic(err)
		} else {
			hex := HashDigit(sha512.Sum512(bs))
			hstr := hex.String()
			specs = append(specs, FileSpec{path, size, mtime, hstr})
		}
	}
	return specs
}
