package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Обработчик возврашает все имеющиеся логи
func (r *Router) getLogs(c *gin.Context) {
	log, _ := r.service.Logs.GetLogs()
	c.JSON(http.StatusOK, gin.H{"logs": log})
}
