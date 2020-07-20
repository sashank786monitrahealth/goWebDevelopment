package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	var f *os.File
	var err error

	f, err = os.Open("toby.jpg")

	if err != nil {
		http.Error(w, "file not found1.", http.StatusNotFound)
		return
	}

	defer f.Close()

	var fi os.FileInfo
	fi, err = f.Stat()

	if err != nil {
		http.Error(w, "file not found1", 404)
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
