package component

import "context"

type userKey struct{}

type User struct {
	Id       int
	UserId   string
	Username string
	Nickname string
	Status   string
}

func UserFormContext(ctx context.Context) (*User, bool) {
	v, ok := ctx.Value(userKey{}).(*User)
	return v, ok
}

func UserMustFormContext(ctx context.Context) *User {
	v, ok := ctx.Value(userKey{}).(*User)
	if !ok {
		panic("user must from context")
	}
	return v
}

func NewUserContext(ctx context.Context, m *User) context.Context {
	return context.WithValue(ctx, userKey{}, m)
}
