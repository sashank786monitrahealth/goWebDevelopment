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

	var sages map[string]string = map[string]string{"India": "Gandhi", "America": "MLK", "Meditate": "Buddha", "Love": "Jesus", "Prayer": "Mohammad", "Understand": "Osho", "Wisdom": "Shailendra Saraswati"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)

	if err != nil {
		log.Fatalln(err)
	}

}
