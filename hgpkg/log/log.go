package log

import (
	"fmt"
)

// For LogFile

type Logger struct {
	filepath string
	buffer   []string
}

func NewLogger(filepath string) *Logger {
	l := new(Logger)
	l.filepath = filepath
	l.buffer = make([]string, 0, 1024)
	return l
}

func (l *Logger) print(level string, message string, err error) {
	l.buffer = append(l.buffer, fmt.Sprintf("[%s] %s >>> %s", level, message, err))
}

func (l *Logger) flash() {
	// Output buffer to log file.
	l.buffer = nil
}

func (l *Logger) Debug(message string) {
	l.print("DEBUG", message, nil)
}

func (l *Logger) Info(message string) {
	l.print("INFO", message, nil)
}

func (l *Logger) Warn(message string, err *error) {
	l.print("WARN", message, *err)
}

func (l *Logger) Error(message string, err *error) {
	l.print("ERROR", message, *err)
}
