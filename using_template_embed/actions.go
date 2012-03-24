package using_template_embed

import (
	"github.com/sunfmin/learn_go_the_hard_way/using_template_embed/templates"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	templates.Get("Hello").Execute(w, name)
}
