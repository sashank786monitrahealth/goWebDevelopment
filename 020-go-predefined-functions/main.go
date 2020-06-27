package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var xs []string = []string{"zero", "one", "two", "three", ""}

	var err error = tpl.Execute(os.Stdout, xs)

	if err != nil {
		log.Fatalln(err)
	}

}
