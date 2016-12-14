package main

import (
	"html/template"
	"net/http"
	"time"
)

type Page struct {
	Now time.Time
}

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("html/top.html")
		if err != nil {
			panic(err)
		}

		page := Page{time.Now()}
		err = tmpl.Execute(w, page)
		if err != nil {
			panic(err)
		}
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:3000", nil)

}
