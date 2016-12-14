package main

import (
	"fmt"
	// "path/filepath"
	"./wire"
)

func main() {
	fmt.Println("Hello")

	managerPath := "_manager.json"
	// rootDir := "."
	// rootDir := "D:/go"

	conf := wire.NewConfig()

	conf.SaveFile()

	man := wire.NewManager(conf)

	if b, _ := man.LoadFile(managerPath); b {
		fmt.Printf("%s file was loaded, files: %d\n", managerPath, man.GetNumFiles())
	}

	fmt.Println("num files: ", man.GetNumFiles())

	man.Evaluate()
	man.SaveFile(managerPath)
}
