package migrations

import (
	"database/sql"
	"embed"
	"fmt"
)

//go:embed *.sql
var all embed.FS

func Run(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	entries, err := all.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("error reading migrations: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		content, err := all.ReadFile(entry.Name())
		if err != nil {
			return nil, err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
