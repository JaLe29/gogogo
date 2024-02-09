package logger

import "time"

type LogMessage struct {
	Level     Level
	Timestamp time.Time
	LogCaller string
	Message   string
}
