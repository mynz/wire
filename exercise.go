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

	/*
	 *     if false {
	 *         top := make([]wire.FileSpec, 0)
	 *         bs, err := ioutil.ReadFile(statusPath)
	 *         if err != nil {
	 *             panic(err)
	 *         }
	 *
	 *         if err := json.Unmarshal(bs, &top); err != nil {
	 *             fmt.Println(err)
	 *             panic(err)
	 *         }
	 *
	 *         fmt.Println(top)
	 *     }
	 *
	 *     if false {
	 *         fileSpecs := wire.CollectFileSpecs(rootDir)
	 *         fmt.Println(fileSpecs)
	 *
	 *         bs, _ := json.MarshalIndent(fileSpecs, "", "  ")
	 *         fmt.Println(string(bs[:]))
	 *
	 *         if err := ioutil.WriteFile(statusPath, bs, 0666); err != nil {
	 *             panic(err)
	 *         }
	 *     }
	 */
}
