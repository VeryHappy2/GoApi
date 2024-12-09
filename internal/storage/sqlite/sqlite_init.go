package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/sqlite"
)

type Storage struct {
	DB *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const operation = "storage.postgres.New"

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &Storage{DB: db}, nil
}
