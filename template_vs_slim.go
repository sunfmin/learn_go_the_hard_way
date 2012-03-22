package main

import (
	"os"
	"text/template"
)

type Data struct {
	Header string
	Items  []*Item
}

type Item struct {
	Name    string
	Current bool
	Url     string
}

func main() {
	data := &Data{
		Header: "Colors",
		Items: []*Item{
			{"red", true, "#red"},
			{"green", false, "#green"},
			{"blue", false, "#blue"},
		},
	}
	tmpl, _ := template.ParseFiles("templates/slim.html")

	for i := 0; i < 30000; i++ {
		tmpl.Execute(os.Stdout, data)
	}
}
