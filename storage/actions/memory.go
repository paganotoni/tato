package actions

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/paganotoni/tato"
)

type Memory struct {
	Actions []tato.Action
}

func (ss *Memory) Create(ctx context.Context, ac tato.Action) error {
	ss.Actions = append(ss.Actions, ac)
	return nil
}

func (ss *Memory) List(ctx context.Context) ([]tato.Action, error) {
	return ss.Actions, nil
}

func (ss *Memory) Destroy(ctx context.Context, id string) error {
	index := -1
	for i, v := range ss.Actions {
		if v.ID != id {
			continue
		}

		index = i
		break
	}

	if index == -1 {
		return sql.ErrNoRows
	}

	ss.Actions = append(ss.Actions[:index], ss.Actions[index+1:]...)

	return nil
}
