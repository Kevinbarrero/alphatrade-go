package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kevinbarrero/alphatrade-go/binance"
	"github.com/kevinbarrero/alphatrade-go/util"
	_ "github.com/lib/pq"
)

type GetCoinRequest struct {
	Symbol    string `json:"symbol"`
	Kline     string `json:"kline"`
	OpenTime  uint64 `json:"open_time"`
	CloseTime uint64 `json:"close_time"`
}
type GetcoinResponse struct {
	Symbol string `json:"symbol"`
	Klines Kline  `json:"klines"`
}

type Kline struct {
	OpenTime  uint64  `json:"open_time"`
	Open      float32 `json:"open"`
	High      float32 `json:"high"`
	Low       float32 `json:"low"`
	Close     float32 `json:"close"`
	Volume    float32 `json:"volume"`
	CloseTime uint64  `json:"close_time"`
	NTrades   uint64  `json:"n_trades"`
}

func GetCoin(db *sql.DB, tableName string) {
	query := fmt.Sprintf(`
		SELECT * FROM %s

	`)
}
func UpdateCreateCoin(db *sql.DB, tableName string) error {
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

// has to be a routine
func UpdateCoins() {
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
		if err := UpdateCreateCoin(db, i.Symbol); err != nil {
			log.Fatal(err)
		}
	}

}
