package web

import (
	"net/http"

	"github.com/paganotoni/tato/web/api/actions"
	"github.com/paganotoni/tato/web/assets"
	"github.com/paganotoni/tato/web/game"
)

func Server() http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("/api/actions/create", actions.Create) //TODO: Only POST!
	server.HandleFunc("/api/actions/list", actions.List)
	server.HandleFunc("/api/actions/destroy/", actions.Destroy)

	server.HandleFunc("/api/stats/k1-distribution", actions.Distribution)

	server.Handle("/static/", http.StripPrefix("/static/", assets.Server))
	server.HandleFunc("/", game.Handler)

	return server
}
