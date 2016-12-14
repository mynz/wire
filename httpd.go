package main

import (
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Now   time.Time
	Items []string
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
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:3000", nil)
}
