package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kevinbarrero/alphatrade-go/util"
)

func main() {
	Getcoins()
}

func Getcoins() {
	requestURL := util.Config.BinanceURL + "v1/ticker/price"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Errorf("error getting coins", err)
	}
	defer res.Body.Close()
	var coins []coin
	io, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(io), &coins)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%+v", coins)
	for _, i := range coins {
		fmt.Println(i.Symbol)
	}
}
