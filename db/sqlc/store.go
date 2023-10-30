package db

import (
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func coinQuery(db *sql.DB, tableName string) error {
	query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            open_time TIMESTAMP UNIQUE,
            open FLOAT,
            high FLOAT,
            low FLOAT,
            close FLOAT,
            volume FLOAT,
            close_time TIMESTAMP UNIQUE,
            n_trades INT
        );
    `, tableName)

	_, err := db.Exec(query)
	return err
}
