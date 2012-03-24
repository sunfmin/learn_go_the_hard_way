package demohttptest

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// Path variable names are in the URL.Query() and start with ':'.
	name := r.URL.Query().Get("name")
	io.WriteString(w, "Hello, "+name)
}

func splat(w http.ResponseWriter, r *http.Request) {
	// Path variable names are in the URL.Query() and start with ':'.
	s := r.URL.Query().Get(":splat")
	io.WriteString(w, "Splat: "+s)
}
