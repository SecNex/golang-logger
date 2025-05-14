package main

import "log"

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

func (l *Logger) DEBUG(message string) {
	if l.Level > DEBUG {
		log.Printf("[DEBUG] %s", message)
	}
}

func (l *Logger) INFO(message string) {
	if l.Level > INFO {
		log.Printf("[INFO] %s", message)
	}
}

func (l *Logger) WARN(message string) {
	if l.Level > WARN {
		log.Printf("[WARN] %s", message)
	}
}

func (l *Logger) ERROR(message string) {
	if l.Level > ERROR {
		log.Printf("[ERROR] %s", message)
	}
}
