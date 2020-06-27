package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Datatp struct {
	Words []string
	Lname string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var xs []string = []string{"zero", "one", "two", "three", ""}

	var data Datatp = Datatp{
		Words: xs,
		Lname: "Bhogu",
	}

	var err error = tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatalln(err)
	}

}
