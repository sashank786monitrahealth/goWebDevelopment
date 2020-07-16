package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", nh67)
	http.HandleFunc("/nh67-tadipatri.jpg", nh67Pic)
	http.ListenAndServe(":8080", nil)
}

func nh67(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="nh67-tadipatri.jpg">`)
}

func nh67Pic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "nh67-tadipatri.jpg")
}
