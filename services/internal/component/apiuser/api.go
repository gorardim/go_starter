package apiuser

import (
	"context"

	"app/api/model"
)

type key struct{}

func NewContext(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, key{}, user)
}

func FromContext(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(key{}).(*model.User)
	return user, ok
}

func MustFromContext(ctx context.Context) *model.User {
	return ctx.Value(key{}).(*model.User)
}
