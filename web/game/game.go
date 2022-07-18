package game

import (
	_ "embed"
	"net/http"

	"github.com/paganotoni/tato/web/render"
)

//go:embed game.html
var html []byte

func Handler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Game string
	}{
		Game: "ATL vs RUS",
	}

	render.WithLayout(w, data, string(html))
}
