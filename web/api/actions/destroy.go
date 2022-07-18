package actions

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/paganotoni/tato/storage/actions"
)

func Destroy(w http.ResponseWriter, r *http.Request) {

	id := strings.Split(r.URL.Path, "/")[4]
	err := actions.Storage.Destroy(r.Context(), id)

	fmt.Println("ID: ", id)

	if errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
