package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/paganotoni/tato"
	"github.com/paganotoni/tato/storage/actions"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	action := struct {
		Input string `json:"input"`
	}{}

	err = json.Unmarshal(data, &action)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	a, err := tato.Parse(action.Input)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	err = actions.Storage.Create(r.Context(), *a)
	if err != nil {
		fmt.Println("Error creating the action: %w", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
