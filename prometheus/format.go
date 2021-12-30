package prometheus

import (
	"fmt"
)

func CreateMetricString(currencyPair string, value string) string {
	tags := fmt.Sprintf("{currencyPair=\"%s\", }", currencyPair)
	metric := fmt.Sprintf("# TYPE ticker gauge\nticker%s %s", tags, value)
	return metric
}
