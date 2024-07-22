package schemas

import "time"

// Logs представляет собой структуру для хранения логов с информацией о времени, названии приложения и сообщении.
type Logs struct {
	Date    time.Time
	AppName string
	Message string
}
