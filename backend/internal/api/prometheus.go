package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler возвращает HTTP-обработчик для метрик Prometheus
func prometheusHandler() http.Handler {
	return promhttp.Handler()
}
