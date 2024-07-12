package api

import (
	"logcollector/internal/monitoring"
	"logcollector/internal/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	service *service.Service
}

func NewRouter(service *service.Service) *Router {
	return &Router{service: service}
}

// InitRoutes инициализирует маршруты для HTTP-сервера.
func (r *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	// Добавление middleware для регистрации метрик
	router.Use(monitoring.MetricsMiddleware())

	metrics := router.Group("/metrics")
	{
		metrics.GET("/", gin.WrapH(prometheusHandler()))
	}

	logs := router.Group("/logs")
	{
		logs.GET("/", r.getLogs)
	}

	return router
}
