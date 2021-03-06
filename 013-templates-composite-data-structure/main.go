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

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var gandhi sage = sage{Name: "Gandhi",
		Motto: "Be the change you wish to see in the world.",
	}

	var buddha sage = sage{Name: "Buddha",
		Motto: "The belief of no beliefs",
	}

	var jesus sage = sage{Name: "Jesus",
		Motto: "Love all",
	}

	var mlk sage = sage{Name: "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone.",
	}

	var Mohammed sage = sage{Name: "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone.",
	}

	var f car = car{
		Manufacturer: "Ford",
		Model:        "XS-50",
		Doors:        5,
	}

	var t car = car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	var sages []sage = []sage{gandhi, jesus, mlk, Mohammed, buddha}
	var cars []car = []car{f, t}

	var data items = items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}

}
