package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/nh67-tadipatri.jpg", nh67Pic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	   <img src="/nh67-tadipatri.jpg">
	`)
}

func nh67Pic(w http.ResponseWriter, req *http.Request) {
	//var f *os.File
	//var err error
	f, err := os.Open("nh67-tadipatri.jpg")

	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	defer f.Close()

	io.Copy(w, f)
}
