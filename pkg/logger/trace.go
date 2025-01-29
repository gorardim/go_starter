package logger

import (
	"context"
	"github.com/google/uuid"
	"strings"
)

type loggerTraceCtx struct{}

func TraceIdFromLogger(ctx context.Context) string {
	if v := ctx.Value(loggerTraceCtx{}); v != nil {
		return v.(string)
	}
	return genTraceId()
}

func NewLoggerContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerTraceCtx{}, genTraceId())
}

func genTraceId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func NewLoggerContextWithTraceId(ctx context.Context, traceId string) context.Context {
	if traceId == "" {
		traceId = genTraceId()
	}
	return context.WithValue(ctx, loggerTraceCtx{}, traceId)
}

func TraceIdAppendSpanId(traceId string) string {
	return traceId + ":" + genTraceId()[16:]
}
