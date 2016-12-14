package main

import (
	"./wire"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello")

	rootDir := "."
	// rootDir := "D:/go"

	fileSpecs := wire.CollectFileSpecs(rootDir)
	fmt.Println(fileSpecs)

	// json.Marshal(fileSpecs)
	bs, _ := json.MarshalIndent(fileSpecs, "", "  ")
	fmt.Println(string(bs[:]))

	statusPath := "_status.json"
	if err := ioutil.WriteFile(statusPath, bs, 0666); err != nil {
		panic(err)
	}
}
