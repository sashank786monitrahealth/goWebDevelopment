package main

import (
	"html/template"
	"log"
	"os"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var buddha sage = sage{Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", buddha)

	if err != nil {
		log.Fatalln(err)
	}

}
