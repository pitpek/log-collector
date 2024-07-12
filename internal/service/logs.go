package service

import (
	"logcollector/internal/repository"
	"logcollector/internal/schemas"
)

type LogsService struct {
	repo repository.Logs
}

func NewLogsService(repo repository.Logs) *LogsService {
	return &LogsService{repo: repo}
}

func (ls *LogsService) GetLogs() ([]schemas.Logs, error) {
	return ls.repo.GetLogs()
}
