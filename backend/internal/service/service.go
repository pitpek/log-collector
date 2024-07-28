package service

import (
	"logcollector/internal/repository"
	"logcollector/internal/schemas"
)

// Logs представляет интерфейс для получения логов.
type Logs interface {
	AddLog(log schemas.Logs) error
	GetLogs() ([]schemas.Logs, error)
}

// Service представляет собой структуру, объединяющую различные сервисы.
type Service struct {
	Logs
}

// NewService создает новый экземпляр Service с предоставленным репозиторием.
// repo - репозиторий, который будет использоваться сервисом.
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Logs: NewLogsService(repo.Logs),
	}
}
