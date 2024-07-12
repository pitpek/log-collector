package service

import (
	"logcollector/internal/repository"
	"logcollector/internal/schemas"
)

type Logs interface {
	GetLogs() ([]schemas.Logs, error)
}

type Service struct {
	Logs
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Logs: NewLogsService(repo.Logs),
	}
}
