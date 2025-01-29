package gormx

import (
	"context"
	"time"

	"app/pkg/logger"
	gLog "gorm.io/gorm/logger"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) LogMode(level gLog.LogLevel) gLog.Interface {
	return l
}

func (l *Logger) Info(ctx context.Context, format string, args ...interface{}) {
	logger.Printf(ctx, format, args...)
}

func (l *Logger) Warn(ctx context.Context, format string, args ...interface{}) {
	logger.Printf(ctx, format, args...)
}

func (l *Logger) Error(ctx context.Context, format string, args ...interface{}) {
	logger.Printf(ctx, format, args...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, affected := fc()
	logger.Printf(ctx, "sql: %s, affected: %d, err: %v", sql, affected, err)
}
