package templates

import (
	"fmt"
	"html/template"
)

var Map map[string]*template.Template

func init() {
	Map = make(map[string]*template.Template)
	Put("Hello", HelloTemplate)
}

func Put(name, tmpl string) {
	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		panic(fmt.Sprintf("Load HelloTemplate Error: %v", err))
	}
	Map[name] = t
}

func Get(name string) *template.Template {
	return Map[name]
}
