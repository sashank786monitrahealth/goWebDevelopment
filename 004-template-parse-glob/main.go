package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)

	err = tpl.ExecuteTemplate(os.Stdout, "four.gohtml", nil)

	//fmt.Printf("error = %v\n", err)

}
