package using_template_embed

import (
	"net/http"
	"using_template_embed/templates"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":name")
	templates.Get("Hello").Execute(w, name)
}
