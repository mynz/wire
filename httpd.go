package main

import (
	"./wire"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Now   time.Time
	Items []string
	// Man   wire.Manager
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("html/top.html")
	if err != nil {
		panic(err)
	}

	page := Page{time.Now(), []string{"foo", "bar", "baz"}}
	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	rootDir := "."

	man := wire.NewManager(rootDir)
	if b, err := man.LoadFile(managerPath); b {
		fmt.Printf("%s file was loaded, files: %d\n", managerPath, man.GetNumFiles())
	} else {
		fmt.Println("could not load", b, err)
	}
	fmt.Println("num files: ", man.GetNumFiles())

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:3000", nil)
}
