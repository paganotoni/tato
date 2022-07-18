package actions

import (
	"encoding/json"
	"net/http"

	"github.com/paganotoni/tato/storage/actions"
)

// List the recoded actions.
func List(w http.ResponseWriter, r *http.Request) {
	type action struct {
		ID        string
		Full      string `json:"full"`
		Timestamp int    `json:"timestamp"`
	}

	found, err := actions.Storage.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var result []action
	for _, v := range found {
		result = append(result, action{
			ID:        v.ID,
			Full:      v.Full(),
			Timestamp: int(v.CreatedAt.Unix()),
		})
	}

	payload, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
