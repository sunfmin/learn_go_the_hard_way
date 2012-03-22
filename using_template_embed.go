package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"using_template_embed"
)

func main() {
	m := pat.New()
	m.Get("/hello/:name", http.HandlerFunc(using_template_embed.Hello))
	http.ListenAndServe("localhost:5000", m)
}
