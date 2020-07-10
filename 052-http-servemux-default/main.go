package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func c(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "cat cat kitty")
}

func main() {

	//var mux *http.ServeMux = http.NewServeMux()
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
