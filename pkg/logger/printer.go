package logger

import (
	"encoding/json"
	"fmt"
)

type FormatPrinter = func(logMessage LogMessage) (string, error)

func ConsolePrinter(logMessage LogMessage) (string, error) {
	return fmt.Sprintf("%s	%s	%s	%s", logMessage.Timestamp.Format("2006-01-02T15:04:05.000Z0700"), logMessage.Level.String(), logMessage.LogCaller, logMessage.Message), nil
}

func JsonPrinter(logMessage LogMessage) (string, error) {
	type jsonLogMessage struct {
		Level     string  `json:"level"`
		Timestamp float64 `json:"timestamp"`
		LogCaller string  `json:"log_caller"`
		Message   string  `json:"message"`
	}

	data, err := json.Marshal(jsonLogMessage{
		Level:     logMessage.Level.String(),
		Timestamp: float64(logMessage.Timestamp.UnixNano()) / 1e9,
		LogCaller: logMessage.LogCaller,
		Message:   logMessage.Message,
	})

	if err != nil {
		return "", err
	}

	return string(data), nil
}
