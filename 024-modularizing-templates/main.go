package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	var err error = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)

	if err != nil {
		log.Fatalln(err)
	}
}
