package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var statement = "42"
	if len(os.Args) > 1 {
		statement = os.Args[1]
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", statement)
	if err != nil {
		log.Fatalln(err)
	}

}
