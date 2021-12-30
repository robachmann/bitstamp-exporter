package service

import (
	"fmt"
	"log"
	"os"
	"robachmann/bitstamp-exporter/bitstamp"
	"strings"
)

type TickerMetric struct {
	CurrencyPair string `json:"currencyPair"`
	Value        string `json:"value"`
}

var defaultCurrencyPairs = initCurrencyPairs()

func initCurrencyPairs() []string {
	defaultCurrencyPairsString := os.Getenv("CURRENCY_PAIRS")
	var currencyPairs []string
	for _, currencyPair := range strings.Split(defaultCurrencyPairsString, ",") {
		if len(currencyPair) > 0 {
			currencyPairs = append(currencyPairs, strings.ToLower(currencyPair))
		}
	}

	if len(currencyPairs) == 0 {
		currencyPairs = append(currencyPairs, "btceur")
	}

	log.Println(fmt.Sprintf("Default currency pairs: %s", currencyPairs))
	return currencyPairs
}

func GetTickerMetric(currencyPair string) TickerMetric {
	ticker := bitstamp.GetTicker(currencyPair)
	metric := TickerMetric{CurrencyPair: currencyPair, Value: ticker.Last}
	return metric
}

func GetTickerMetrics() []TickerMetric {
	var metrics []TickerMetric
	for _, currencyPair := range defaultCurrencyPairs {
		metric := GetTickerMetric(currencyPair)
		metrics = append(metrics, metric)
	}
	return metrics
}
