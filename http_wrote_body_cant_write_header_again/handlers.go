package main

import (
	"code.google.com/p/gorilla/sessions"
	"fmt"
	"net/http"
	"reflect"
)

var store = sessions.NewCookieStore([]byte("d908b1c425062e95d30b8d30de7360c1"))

func UpdateSession1(w http.ResponseWriter, r *http.Request) {
	s, err := store.Get(r, "tricky_session_name")
	if err != nil {
		panic(err)
	}
	s.Values["update1"] = "value1"
	s.Save(r, w)
	render(w, r)
}

func UpdateSession2(session map[interface{}]interface{}, w http.ResponseWriter, r *http.Request) {
	session["update2"] = "value2"
	render(w, r)
}

func Show(w http.ResponseWriter, r *http.Request) {
	render(w, r)
}

func render(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	s1, _ := store.Get(r, "tricky_session_name")

	fmt.Fprintf(w, "%+v", s1.Values)

	fmt.Fprintf(w, `<ul>`)
	fmt.Fprintf(w, `<li><a href="/update1">update1</a></li>`)
	fmt.Fprintf(w, `<li><a href="/update2">update2</a></li>`)
	fmt.Fprintf(w, `<li><a href="/show">show</a></li>`)
	fmt.Fprintf(w, `</ul>`)
}

type SessionHandlerFunc func(session map[interface{}]interface{}, w http.ResponseWriter, r *http.Request)

func WithSession(shf SessionHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := store.Get(r, "tricky_session_name")
		if err != nil {
			panic(err)
		}
		shf(s.Values, w, r)
		s.Save(r, w)
	}
}

func Mux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/update1", UpdateSession1)
	m.HandleFunc("/show", Show)
	m.HandleFunc("/update2", WithSession(UpdateSession2))
	return m
}

func main() {
	http.ListenAndServe("localhost:5000", Mux())
}
