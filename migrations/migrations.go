package migrations

import (
	"database/sql"
	"embed"
	"fmt"
)

//go:embed *.sql
var all embed.FS

func Run(db *sql.DB) error {
	entries, err := all.ReadDir(".")
	if err != nil {
		return fmt.Errorf("error reading migrations: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		content, err := all.ReadFile(entry.Name())
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return err
		}
	}

	return nil
}
