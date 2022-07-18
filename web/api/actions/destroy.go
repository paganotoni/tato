package actions

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paganotoni/tato/storage/actions"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := actions.Storage.Destroy(r.Context(), id)

	fmt.Println(id)

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
