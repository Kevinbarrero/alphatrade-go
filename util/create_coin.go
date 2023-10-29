package util

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

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

func createCoin(coinName string) {
	db, err := sql.Open("postgres", "user=root dbname=alphatrade-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := coinQuery(db, coinName); err != nil {
		log.Fatal(err)
	}
}
