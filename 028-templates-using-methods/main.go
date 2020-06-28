package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var p person = person{
		Name: "Ian Fleming",
		Age:  56,
	}

	var err error = tpl.Execute(os.Stdout, p)

	if err != nil {
		log.Fatalln(err)
	}

}

// Generally speaking, best practice:
// call functions in templates for formatting only; not processing or data

// The main reason why you don't want to do processing in your templates:
// 1. separation of concerns
// 2. if you are using a function more than once in a template
// the server needs to do processing more than once.
// (though standard library might cache processing)
// I've yet to dig into this - if you find the answer, let me know
