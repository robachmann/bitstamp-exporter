#!/usr/bin/env bash
env GOOS=linux GOARCH=arm64 go build .
docker build --platform linux/arm64 -t robachmann/bitstamp-exporter:1.0.0-arm64 .
docker push robachmann/bitstamp-exporter:1.0.0-arm64