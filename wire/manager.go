package wire

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type hashDigit [64]byte

func (hash *hashDigit) String() string {
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

type Manager struct {
	RootDir   string     `json:"rootDir"`
	FileSpecs []FileSpec `json:"fileSpecs"`
}

func NewManager(rootDir string) *Manager {
	man := new(Manager)
	man.RootDir = rootDir
	return man
}

func (man *Manager) GetNumFiles() int {
	return len(man.FileSpecs)
}

func collectFileSpecs(rootDir string) []FileSpec {
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
			hex := hashDigit(sha512.Sum512(bs))
			hstr := hex.String()
			specs = append(specs, FileSpec{path, size, mtime, hstr})
		}
	}
	return specs
}

func (man *Manager) Evaluate() {
	man.FileSpecs = collectFileSpecs(man.RootDir)
}

func (man *Manager) SaveFile(path string) error {
	bs, _ := json.MarshalIndent(man, "", "  ")
	if err := ioutil.WriteFile(path, bs, 0666); err != nil {
		return err
	}
	return nil
}

func (man *Manager) LoadFile(path string) (bool, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(bs, &man); err != nil {
		return false, err
	}
	return true, nil
}
