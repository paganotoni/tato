package web

import (
	"net/http"

	"github.com/paganotoni/tato/web/api/actions"
	"github.com/paganotoni/tato/web/assets"
	"github.com/paganotoni/tato/web/game"
	"github.com/paganotoni/tato/web/home"

	"github.com/gorilla/mux"
)

func Server() http.Handler {
	server := mux.NewRouter()

	server.HandleFunc("/", home.Home)
	server.HandleFunc("/game/{id}", game.Handler).Methods(http.MethodGet)

	api := server.PathPrefix("/api/{game_id}/").Subrouter()
	api.HandleFunc("/actions/create", actions.Create).Methods(http.MethodPost)
	api.HandleFunc("/actions/list", actions.List).Methods(http.MethodGet)
	api.HandleFunc("/actions/destroy/{id}", actions.Destroy).Methods(http.MethodDelete)
	api.HandleFunc("/stats/k1", actions.Distribution).Methods(http.MethodGet)

	staticDir := "/static/"
	server.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, assets.Server))

	return server
}
