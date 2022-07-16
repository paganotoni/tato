package actions

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/paganotoni/tato"
)

func List(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type action struct {
			Full      string `json:"full"`
			Timestamp string `json:"timestamp"`
		}

		acs := []tato.Action{}
		result := []action{}
		for _, v := range acs {
			result = append(result, action{
				Full:      v.Full(),
				Timestamp: v.Timestamp.Format("2006-01-02 15:04:05"),
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
}
