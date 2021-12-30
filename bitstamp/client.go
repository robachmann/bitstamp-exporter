package bitstamp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Ticker struct {
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	High      string `json:"high"`
	Last      string `json:"last"`
	Low       string `json:"low"`
	Open      string `json:"open"`
	Timestamp string `json:"timestamp"`
	Volume    string `json:"volume"`
	Vwap      string `json:"vwap"`
}

func GetTicker(currencyPair string) Ticker {
	uri := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/%s/", currencyPair)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	ticker := Ticker{}
	err = json.Unmarshal(body, &ticker)
	if err != nil {
		return Ticker{}
	}
	return ticker
}
