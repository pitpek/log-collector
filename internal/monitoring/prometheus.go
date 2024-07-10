package monitoring

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)
	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func init() {
	// Регистрация метрик в Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
}

// MetricsMiddleware регистрирует метрики для каждого HTTP-запроса
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Обработка запроса
		c.Next()

		duration := time.Since(start).Seconds()
		path := c.FullPath()
		method := c.Request.Method

		// Регистрация метрик
		httpRequestsTotal.WithLabelValues(path, method).Inc()
		httpRequestDuration.WithLabelValues(path, method).Observe(duration)
	}
}
