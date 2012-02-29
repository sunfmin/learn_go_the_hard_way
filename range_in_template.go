package main

import (
	"os"
	"text/template"
)

type Page struct {
	Orders []*Order
}

type Order struct {
	ImageUrl string
}

func main() {
	p := &Page{}
	p.Orders = []*Order{
		{"url1"},
		{"url2"},
		{"url3"},
		{"url4"},
		{"url5"},
	}

	t, _ := template.New("").Parse(`
{{range .Orders}}  
	{{.ImageUrl}}
{{end}}`)
	t.Execute(os.Stdout, p)
}
