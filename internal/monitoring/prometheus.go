package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func init() {
	// Регистрация метрик в Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

// RecordRequest увеличивает счетчик HTTP-запросов
func RecordRequest(path string) {
	httpRequestsTotal.WithLabelValues(path).Inc()
}

// NewRequestTimer создает новый таймер для измерения времени выполнения запроса
func NewRequestTimer(path string) *prometheus.Timer {
	return prometheus.NewTimer(httpRequestDuration.WithLabelValues(path))
}

// PrometheusHandler возвращает HTTP-обработчик для метрик Prometheus
func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
