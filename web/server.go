package web

import (
	"database/sql"
	"net/http"

	"github.com/paganotoni/tato/web/api/actions"
	"github.com/paganotoni/tato/web/assets"
	"github.com/paganotoni/tato/web/game"
)

func Server(db *sql.DB) http.Handler {
	server := http.NewServeMux()

	server.HandleFunc("/api/actions/create", actions.Create(db)) //TODO: Only POST!
	server.HandleFunc("/api/actions/list", actions.List(db))

	server.Handle("/static/", http.StripPrefix("/static/", assets.Server))
	server.HandleFunc("/", game.Handler)

	return server
}
