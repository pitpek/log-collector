package schemas

import "time"

type Logs struct {
	Date    time.Time
	AppName string
	Message string
}
