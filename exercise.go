package main

import (
	"fmt"
	// "path/filepath"
	"./wire"
)

func main() {
	fmt.Println("Hello")

	conf := wire.NewConfig()
	if conf.LoadFile() == nil {
		fmt.Println("config file was loaded")
	}
	conf.SaveFile()

	man := wire.NewManager(conf)

	if b, _ := man.LoadFile(wire.ManagerPath); b {
		fmt.Printf("%s file was loaded, files: %d\n", wire.ManagerPath, man.GetNumFiles())
	}

	fmt.Println("num files: ", man.GetNumFiles())

	man.Evaluate()
	man.SaveFile(wire.ManagerPath)
}
