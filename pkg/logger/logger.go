package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type logger struct {
	Level            Level
	FormatPrinter    FormatPrinter
	workingDirectory string
}

func (l *logger) printMessage(level Level, msg string) {
	if l.Level > level {
		return
	}

	logMessage := LogMessage{
		Level:     level,
		Timestamp: time.Now(),
		LogCaller: l.getCaller(),
		Message:   msg,
	}

	str, err := l.FormatPrinter(logMessage)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(str)
}

func (l *logger) getCaller() string {
	_, filename, no, ok := runtime.Caller(3)
	if !ok {
		fmt.Println("Failed to get caller information")
		return ""
	}

	relFileName := strings.TrimPrefix(filename, l.workingDirectory)

	return fmt.Sprintf("%s:%d", relFileName, no)
}

func (l *logger) Debug(msg string) {
	l.printMessage(LevelDebug, msg)
}

func (l *logger) Info(msg string) {
	l.printMessage(LevelInfo, msg)
}

func (l *logger) Warn(msg string) {
	l.printMessage(LevelWarn, msg)
}

func (l *logger) Error(msg string) {
	l.printMessage(LevelError, msg)
}

func New() Logger {
	wd, err := os.Getwd()
	if err != nil {
		panic("Failed to get working directory")
	}

	return &logger{
		Level:            0,
		FormatPrinter:    ConsolePrinter,
		workingDirectory: wd + string(os.PathSeparator),
	}
}
