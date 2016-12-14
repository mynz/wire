package main

import (
	"./wire"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	statusPath := "_status.json"
	rootDir := "."
	// rootDir := "D:/go"

	man := wire.NewManager(rootDir)

	if b, _ := man.LoadFile(statusPath); b {
		fmt.Printf("%s file was loaded, files: %d\n", statusPath, man.GetNumFiles())
	}

	man.Evaluate()
	man.SaveFile(statusPath)
}
