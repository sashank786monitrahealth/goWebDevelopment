package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var sages []string = []string{"Gandhi", "MLK", "Buddha", "Jesus", "Mohammad", "Osho", "Shailendra Saraswati"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)

	if err != nil {
		log.Fatalln(err)
	}

}
