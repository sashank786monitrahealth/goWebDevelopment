package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "cat cat kitty")
}

func main() {
	var d hotdog
	var c hotcat

	var mux *http.ServeMux = http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
