package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"robachmann/bitstamp-exporter/api"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", api.HomeLink)
	router.HandleFunc("/ticker/{currencyPair}", api.GetTickerByCurrencyPair).Methods("GET")
	router.HandleFunc("/metrics", api.GetMetrics).Methods("GET")                              //.Headers("Accept", "plain/text;version=0.0.4")
	router.HandleFunc("/metrics/{currencyPair}", api.GetMetricsByCurrencyPair).Methods("GET") //.Headers("Accept", "plain/text;version=0.0.4")
	log.Fatal(http.ListenAndServe(":8080", router))
}
