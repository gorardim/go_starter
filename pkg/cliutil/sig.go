package cliutil

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

var showdownFuncList = make([]func(), 0)

func RegisterShowdown(fn func()) {
	showdownFuncList = append(showdownFuncList, fn)
}

func NewShutDown() (context.Context, func()) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return ctx, func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		cancelFunc()
		for _, f := range showdownFuncList {
			f()
		}
	}
}
