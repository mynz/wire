package main

import (
	"./wire"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	managerPath := "_manager.json"
	rootDir := "."
	// rootDir := "D:/go"

	man := wire.NewManager(rootDir)

	if b, _ := man.LoadFile(managerPath); b {
		fmt.Printf("%s file was loaded, files: %d\n", managerPath, man.GetNumFiles())
	}

	man.Evaluate()
	man.SaveFile(managerPath)
}
