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

	var sages []sage = []sage{gandhi, jesus, mlk, Mohammed, buddha}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)

	if err != nil {
		log.Fatalln(err)
	}

}
