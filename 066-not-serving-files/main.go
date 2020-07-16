package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--Not Serving from our server-->
	  <img src="https://upload.wikimedia.org/wikipedia/commons/b/ba/Daily-Login-Sashank-2020-June-25.jpg">
	  `,
	)
}
