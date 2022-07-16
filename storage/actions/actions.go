package actions

import (
	"context"
	"database/sql"

	"github.com/paganotoni/tato"
)

type storage interface {
	Create(context.Context, tato.Action) error
	List(context.Context) []tato.Action
}

func NewSQLiteStorage(db *sql.DB) storage {
	return &sqliteStorage{db: db}
}
