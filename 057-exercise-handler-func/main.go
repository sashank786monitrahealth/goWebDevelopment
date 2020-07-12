/*
1. Take the previous program in the previous folder and change it so that:
   a) a template is parsed and served
   b) you pass data into the template
*/

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(bhogu))
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func bar(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "bar ran")
}

func bhogu(res http.ResponseWriter, req *http.Request) {
	var tpl *template.Template
	var err error
	tpl, err = template.ParseFiles("something.gohtml")

	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(res, "something.gohtml", "Bhogu")

	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
