package actions

import (
	"context"

	"github.com/paganotoni/tato"
)

// Shared service implementation, can be overriden when
// testing.
var Storage service = &sqliteStorage{}

type service interface {
	Create(context.Context, tato.Action) error
	Destroy(context.Context, string) error
	List(context.Context) ([]tato.Action, error)
}
