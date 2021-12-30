package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"robachmann/bitstamp-exporter/prometheus"
	"robachmann/bitstamp-exporter/service"
	"strings"
)

func HomeLink(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome home!")
	if err != nil {
		log.Println(err)
	}
}

func GetTickerByCurrencyPair(writer http.ResponseWriter, request *http.Request) {
	currencyPair := mux.Vars(request)["currencyPair"]

	ticker := service.GetTickerMetric(currencyPair)
	result, err := json.Marshal(ticker)
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(result)
}

func GetMetrics(writer http.ResponseWriter, _ *http.Request) {

	var tickerMetrics []service.TickerMetric
	tickerMetrics = service.GetTickerMetrics()

	var results []string
	for _, metric := range tickerMetrics {
		tickerString := prometheus.CreateMetricString(metric.CurrencyPair, metric.Value)
		results = append(results, tickerString)
	}

	result := strings.Join(results[:], "\n")

	writer.Header().Set("Content-Type", "plain/text;version=0.0.4")
	_, _ = writer.Write([]byte(result))
}

func GetMetricsByCurrencyPair(writer http.ResponseWriter, request *http.Request) {
	currencyPair := mux.Vars(request)["currencyPair"]
	metric := service.GetTickerMetric(currencyPair)
	result := prometheus.CreateMetricString(metric.CurrencyPair, metric.Value)
	writer.Header().Set("Content-Type", "plain/text;version=0.0.4")
	_, _ = writer.Write([]byte(result))
}
