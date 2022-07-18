package home

import (
	_ "embed"
	"net/http"

	"github.com/paganotoni/tato/web/render"
)

//go:embed home.html
var html string

func Home(w http.ResponseWriter, r *http.Request) {
	render.WithLayout(w, nil, string(html))
}
