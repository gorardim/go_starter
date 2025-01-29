package logger

import (
	"app/pkg/utils"
	"context"
	"fmt"
	"log"
	"os"
)

var _ Logger = (*StdLogger)(nil)

type StdLogger struct {
	*log.Logger
}

func (s *StdLogger) PrintMap(ctx context.Context, m map[string]string) {
	s.Output(3, fmt.Sprintln(TraceIdFromLogger(ctx), utils.JsonEncode(m)))
}

func NewStdLogger() *StdLogger {
	return &StdLogger{Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)}
}

func (s *StdLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	s.Output(3, fmt.Sprintln(TraceIdFromLogger(ctx), fmt.Sprintf(format, v...)))
}

func (s *StdLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	s.Output(3, fmt.Sprintln(TraceIdFromLogger(ctx), fmt.Sprintf(format, v...)))
	os.Exit(1)
}
