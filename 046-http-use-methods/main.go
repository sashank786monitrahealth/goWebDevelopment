package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

type inputData struct {
	Method      string
	Submissions url.Values
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error = req.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	var data inputData = inputData{
		Method:      req.Method,
		Submissions: req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
