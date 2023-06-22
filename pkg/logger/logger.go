// Package logger deals with logging.
package logger

// Logger declares methods for logging.
type Logger interface {
	Info(message ...any)
	Error(message ...any)
}
