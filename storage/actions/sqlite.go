package actions

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/paganotoni/tato"
)

var (
	//go:embed insert.sql
	insertStatement string
	//go:embed list.sql
	listQuery string
)

type sqliteStorage struct {
	db *sql.DB
}

func (ss *sqliteStorage) Create(ctx context.Context, ac tato.Action) error {
	_, err := ss.db.ExecContext(
		ctx,
		insertStatement,

		ac.ID,
		ac.Player,
		ac.Kind,
		ac.Class,
		ac.Evaluation,
		ac.StartingZone,
		ac.EndingZone,
	)

	return err
}

func (ss *sqliteStorage) List(ctx context.Context) []tato.Action {
	var result []tato.Action
	rows, err := ss.db.QueryContext(ctx, listQuery)
	if err != nil {
		return result
	}

	for rows.Next() {
		var ac tato.Action
		err := rows.Scan(
			&ac.ID,
			&ac.Player,
			&ac.Kind,
			&ac.Class,
			&ac.Evaluation,
			&ac.StartingZone,
			&ac.EndingZone,
			&ac.Timestamp,
		)
		if err != nil {
			continue
		}

		result = append(result, ac)
	}

	return result
}
