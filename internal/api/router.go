package api

import (
	"logcollector/internal/monitoring"

	"github.com/gin-gonic/gin"
)

// InitRoutes инициализирует маршруты для HTTP-сервера.
func InitRoutes() *gin.Engine {
	router := gin.New()

	// Добавление middleware для регистрации метрик
	router.Use(monitoring.MetricsMiddleware())

	metrics := router.Group("/metrics")
	{
		metrics.GET("/", gin.WrapH(monitoring.PrometheusHandler()))
	}

	return router
}
