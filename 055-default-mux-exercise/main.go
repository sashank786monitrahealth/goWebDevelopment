package main

import (
	"io"
	"net/http"
)

/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default serve mux

"/"
"/dog/"
"/me/"

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func myName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hey sash")
}
