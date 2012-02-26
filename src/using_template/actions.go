package using_template

import (
	"net/http"
	"using_template/templates"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":name")
	templates.Get("Hello").Execute(w, name)
}
