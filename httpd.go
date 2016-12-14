package main

import (
	"net/http"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:3000", nil)

}
