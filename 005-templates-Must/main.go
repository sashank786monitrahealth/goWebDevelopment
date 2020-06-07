package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	var err error = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)

	err = tpl.ExecuteTemplate(os.Stdout, "four.gohtml", nil)

	//fmt.Printf("error = %v\n", err)

}
