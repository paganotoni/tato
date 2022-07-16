package game

import (
	_ "embed"
	"net/http"
)

//go:embed game.html
var html []byte

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
}
