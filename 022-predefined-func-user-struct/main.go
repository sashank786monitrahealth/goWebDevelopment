package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type User struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var u1 User = User{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	var u2 User = User{
		Name:  "Gandhi",
		Motto: "Be the change you wish to see in this world.",
		Admin: true,
	}

	var u3 User = User{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	var users []User = []User{u1, u2, u3}

	var err error = tpl.Execute(os.Stdout, users)

	if err != nil {
		log.Fatalln(err)
	}

}
