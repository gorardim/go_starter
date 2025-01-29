package logger

import (
	"context"
)

var DefaultLogger Logger = NewStdLogger()

func Printf(ctx context.Context, format string, v ...interface{}) {
	DefaultLogger.Printf(ctx, format, v...)
}

func PrintMap(ctx context.Context, m map[string]string) {
	DefaultLogger.PrintMap(ctx, m)
}

func Fatalf(ctx context.Context, format string, v ...interface{}) {
	DefaultLogger.Fatalf(ctx, format, v...)
}

type Logger interface {
	Printf(ctx context.Context, format string, v ...interface{})
	Fatalf(ctx context.Context, format string, v ...interface{})
	PrintMap(ctx context.Context, m map[string]string)
}
