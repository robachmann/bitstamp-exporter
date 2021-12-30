FROM arm64v8/alpine
COPY ./bitstamp-exporter bitstamp-exporter
ENTRYPOINT ["./bitstamp-exporter"]