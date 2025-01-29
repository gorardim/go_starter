package jwtutil

import (
	"golang.org/x/net/context"
)

type dataKey struct{}

func NewContext[T any](ctx context.Context, data *T) context.Context {
	return context.WithValue(ctx, dataKey{}, data)
}

func FromContext[T any](ctx context.Context) (*T, bool) {
	d, ok := ctx.Value(dataKey{}).(*T)
	return d, ok
}
