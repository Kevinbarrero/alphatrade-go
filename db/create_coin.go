package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kevinbarrero/alphatrade-go/binance"
	"github.com/kevinbarrero/alphatrade-go/util"
	_ "github.com/lib/pq"
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

func createCoin() {
	db, err := sql.Open("postgres", "user=root dbname=alphatrade-go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	coins, err := binance.GetBinanceCoins(config)
	if err != nil {
		log.Printf("cannot get coins from binance %s", err)
	}
	for _, i := range coins {
		if err := coinQuery(db, i.Symbol); err != nil {
			log.Fatal(err)
		}
	}

}
