package main

import "net/http"

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

}

func dogPic(w http.ResponseWriter, req *http.Request) {

}
