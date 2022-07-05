//go:generate mockgen -source logger.go -destination mocks/logger_mock.go -package mocks
package logger

import (
	"context"
)

type Logger interface {
	Error(args ...interface{})
	ErrorContext(ctx context.Context, args ...interface{})
	Info(args ...interface{})
	InfoContext(ctx context.Context, args ...interface{})
	Fatal(args ...interface{})
}
