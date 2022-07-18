package render

import (
	_ "embed"

	"fmt"
	"html/template"
	"net/http"
)

var (
	//go:embed layout.html
	base string

	layout *template.Template
)

func WithLayout(w http.ResponseWriter, data interface{}, content string) {
	if layout == nil {
		var err error
		layout, err = template.New("layout").Parse(base)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tt := fmt.Sprintf(`{{define "content"}}%s{{end}}`, content)
	tmpl, err := template.Must(layout.Clone()).Parse(tt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
