package binance

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kevinbarrero/alphatrade-go/util"
)

type Coin struct {
	Symbol string `json:"symbol"`
}

func GetBinanceCoins(config util.Config) ([]Coin, error) {
	requestURL := config.BinanceURL + "v1/ticker/price"

	res, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var coins []Coin

	io, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(io), &coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}
