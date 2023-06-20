// Package zap implements logging with zap logger.
package zap

import (
	"io"
	"os"

	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct {
	zap *zap.SugaredLogger
}

// New returns a new logger that logs into file and stdout.
func New(w io.Writer, level zapcore.Level) logger.Logger {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	writer := zapcore.AddSync(w)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	logger := zap.New(core)
	sugarLogger := logger.Sugar()

	return &log{zap: sugarLogger}
}

// Info logs an info.
func (l *log) Info(message string) {
	l.zap.Info(message)
}

// Info logs an Error.
func (l *log) Error(message string) {
	l.zap.Error(message)
}

// Sync flushes any buffered log entries.
func (l *log) Sync() error {
	return l.zap.Sync()
}
