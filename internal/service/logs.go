package service

import (
	"logcollector/internal/repository"
	"logcollector/internal/schemas"
	"time"
)

// LogsService представляет собой сервис для работы с логами, использующий репозиторий логов.
type LogsService struct {
	repo repository.Logs
}

// NewLogsService создает новый экземпляр LogsService с предоставленным репозиторием логов.
// repo - репозиторий логов, который будет использоваться сервисом.
func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{repo: repo}
}

// AddLog добавляет новый лог в репозиторий.
// date - дата и время сообщения
// key - название приложения, с которого пришло сообщение
// message - сообщение, которое нужно вставить
func (ls *LogsService) AddLog(date time.Time, key, message string) error {
	return ls.repo.AddLog(date, key, message)
}

// GetLogs извлекает все логи из репозитория.
func (ls *LogsService) GetLogs() ([]schemas.Logs, error) {
	return ls.repo.GetLogs()
}
