package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var y year = year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to Multimedia", "3"},
				course{"CSCI-41", "Introduction to Networking", "4"},
				course{"CSCI-42", "Introduction to Operating Systems", "2"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Introduction to Go", "3"},
				course{"CSCI-51", "Introduction to Networking in Go", "4"},
				course{"CSCI-52", "Introduction to Printing in Go", "2"},
			},
		},
	}

	var err error = tpl.Execute(os.Stdout, y)

	if err != nil {
		log.Fatalln(err)
	}
}
