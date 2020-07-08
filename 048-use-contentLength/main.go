package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

type inputData struct {
	Method        string
	URL           *url.URL
	Submissions   map[string][]string
	Header        http.Header
	Host          string
	ContentLength int64
}

var tpl *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error = req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	var data inputData = inputData{
		Method:        req.Method,
		URL:           req.URL,
		Submissions:   req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)

}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
