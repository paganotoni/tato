package actions

import (
	"encoding/json"
	"net/http"

	"github.com/paganotoni/tato/storage/actions"
)

type distributionResult struct {
	Total   int `json:"total"`
	Perfect int `json:"perfect"`
	Error   int `json:"error"`
}

func Distribution(w http.ResponseWriter, r *http.Request) {
	ac, err := actions.Storage.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	result := map[string]*distributionResult{}
	for _, v := range ac {
		if v.Kind != "A" {
			continue
		}

		zone := result[v.StartingZone]
		if zone == nil {
			zone = &distributionResult{}
		}

		zone.Total++
		if v.Evaluation == "*" {
			zone.Perfect++
		}

		if v.Evaluation == "=" {
			zone.Error++
		}

		result[v.StartingZone] = zone
	}

	payload, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}
