package storage

import (
	"database/sql"
)

// DB shared across the app for SQLite operations.
var DB *sql.DB
