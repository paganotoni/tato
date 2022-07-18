package actions_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paganotoni/tato"
	sta "github.com/paganotoni/tato/storage/actions"
	"github.com/paganotoni/tato/web/api/actions"
)

func TestDestroy(t *testing.T) {

	t.Run("Invalid ID", func(t *testing.T) {
		ac, _ := tato.Parse("12A+")
		mst := &sta.Memory{
			Actions: []tato.Action{*ac},
		}

		sta.Storage = mst

		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/actions/1", nil)

		actions.Destroy(w, r)
		if len(mst.Actions) != 1 {
			t.Errorf("Expected 1 action, got %d", len(mst.Actions))
		}

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("Valid ID", func(t *testing.T) {
		ac, _ := tato.Parse("12A+")
		mst := &sta.Memory{
			Actions: []tato.Action{*ac},
		}

		sta.Storage = mst

		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/actions/"+ac.ID, nil)

		actions.Destroy(w, r)
		if len(mst.Actions) != 0 {
			t.Errorf("Expected 0 action, got %d", len(mst.Actions))
		}
	})

}
