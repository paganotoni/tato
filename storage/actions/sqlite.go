package actions

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/paganotoni/tato"
	"github.com/paganotoni/tato/storage"
)

var (
	//go:embed insert.sql
	insertStatement string
	//go:embed list.sql
	listQuery string
)

type sqliteStorage struct{}

func (ss *sqliteStorage) Create(ctx context.Context, ac tato.Action) error {
	_, err := storage.DB.ExecContext(
		ctx,
		insertStatement,

		ac.ID,
		ac.Player,
		ac.Kind,
		ac.Class,
		ac.Evaluation,
		ac.StartingZone,
		ac.EndingZone,
		ac.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating action: %w", err)
	}

	return nil
}

func (ss *sqliteStorage) List(ctx context.Context) ([]tato.Action, error) {
	var result []tato.Action

	rows, err := storage.DB.QueryContext(ctx, listQuery)
	if err != nil {
		return result, fmt.Errorf("error listing actions: %w", err)
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
			&ac.CreatedAt,
		)

		if err != nil {
			return result, fmt.Errorf("error scanning actions result: %w", err)
		}

		result = append(result, ac)
	}

	return result, nil
}
