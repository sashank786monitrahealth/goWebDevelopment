package main

import (
	"html/template"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters

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
	//tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	tpl = template.New("").Funcs(fm)
	tpl = template.Must(tpl.ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func double(x float64) float64 {
	return x + x
}

func square(x float64) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 36.0)

	if err != nil {
		log.Fatalln(err)
	}

}
