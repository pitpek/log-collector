package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PrometheusHandler возвращает HTTP-обработчик для метрик Prometheus
func prometheusHandler() http.Handler {
	return promhttp.Handler()
}

func (r *Router) getLogs(c *gin.Context) {
	log, _ := r.service.Logs.GetLogs()
	c.JSON(http.StatusOK, gin.H{
		"message": "Logs retrieved successfully", "logs": log,
	})
}
