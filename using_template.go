package main

import (
	"github.com/bmizerany/pat.go"
	"net/http"
	"using_template"
)

func main() {
	m := pat.New()
	m.Get("/hello/:name", http.HandlerFunc(using_template.Hello))
	http.ListenAndServe("localhost:5000", m)
}
