package web

import (
	"net/http"

	"github.com/paganotoni/tato/web/game"
)

func Server() http.Handler {
	server := http.NewServeMux()
	server.HandleFunc("/", game.Handler)

	return server
}
