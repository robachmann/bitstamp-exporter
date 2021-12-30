package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var defaultCurrencyPairs []string

func main() {
	initCurrencyPairs()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/ticker/{currencyPair}", getTickerByCurrencyPair).Methods("GET")
	router.HandleFunc("/metrics", getMetrics).Methods("GET")                              //.Headers("Accept", "plain/text;version=0.0.4")
	router.HandleFunc("/metrics/{currencyPair}", getMetricsByCurrencyPair).Methods("GET") //.Headers("Accept", "plain/text;version=0.0.4")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initCurrencyPairs() {
	defaultCurrencyPairsString := os.Getenv("CURRENCY_PAIRS")

	for _, currencyPair := range strings.Split(defaultCurrencyPairsString, ",") {
		if len(currencyPair) > 0 {
			defaultCurrencyPairs = append(defaultCurrencyPairs, strings.ToLower(currencyPair))
		}
	}

	if len(defaultCurrencyPairs) == 0 {
		defaultCurrencyPairs = append(defaultCurrencyPairs, "btceur")
	}

	log.Println(fmt.Sprintf("Default currency pairs: %s", defaultCurrencyPairs))
}

func homeLink(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome home!")
	if err != nil {
		log.Println(err)
	}
}

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

func getTickerByCurrencyPair(writer http.ResponseWriter, request *http.Request) {
	currencyPair := mux.Vars(request)["currencyPair"]

	ticker := getTicker(currencyPair)
	result, err := json.Marshal(ticker)
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(result)
}

func getMetrics(writer http.ResponseWriter, _ *http.Request) {

	var results []string
	for _, currencyPair := range defaultCurrencyPairs {
		ticker := getTicker(currencyPair)
		tickerString := convertToMetric(ticker, currencyPair)
		results = append(results, tickerString)
	}

	result := strings.Join(results[:], "\n")

	writer.Header().Set("Content-Type", "plain/text;version=0.0.4")
	_, _ = writer.Write([]byte(result))
}

func getMetricsByCurrencyPair(writer http.ResponseWriter, request *http.Request) {
	currencyPair := mux.Vars(request)["currencyPair"]
	ticker := getTicker(currencyPair)
	result := convertToMetric(ticker, currencyPair)
	writer.Header().Set("Content-Type", "plain/text;version=0.0.4")
	_, _ = writer.Write([]byte(result))
}

func convertToMetric(ticker Ticker, currencyPair string) string {
	metric := fmt.Sprintf("# TYPE %s %s\n%s%s %s", "ticker", "gauge", "ticker", fmt.Sprintf("{currencyPair=\"%s\", }", currencyPair), ticker.Last)
	return metric
}

func getTicker(currencyPair string) Ticker {
	uri := fmt.Sprintf("https://www.bitstamp.net/api/v2/ticker/%s/", currencyPair)
	//	log.Println(uri)
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
