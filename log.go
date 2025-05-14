package logger

import (
	"log"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	Level
}

func NewLogger(level Level) *Logger {
	return &Logger{
		Level: level,
	}
}

func (l *Logger) SetLevel(level Level) {
	l.Level = level
}

func (l *Logger) Create(level string, message string) {
	log.Printf("[%s] %s", level, message)
}

func (l *Logger) DEBUG(message string) {
	if l.Level <= DEBUG {
		l.Create("DEBUG", message)
	}
}

func (l *Logger) INFO(message string) {
	if l.Level <= INFO {
		l.Create("INFO", message)
	}
}

func (l *Logger) WARN(message string) {
	if l.Level <= WARN {
		l.Create("WARN", message)
	}
}

func (l *Logger) ERROR(message string) {
	if l.Level <= ERROR {
		l.Create("ERROR", message)
	}
}
