package admin

import (
	"context"
	_ "embed"

	"app/pkg/jwtutil"
)

//go:embed openapi.yml
var Openapi string

// AuthUser 授权用户
type AuthUser struct {
	Id     int    `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

func UserFormContext(ctx context.Context) (*AuthUser, bool) {
	return jwtutil.FromContext[AuthUser](ctx)
}

const (
	ErrInvalidParam    = "ErrInvalidParam"
	ErrUnauthenticated = "ErrUnauthenticated"
	ErrNoPermission    = "ErrNoPermission"
	ErrBusiness        = "ErrBusiness"
)

type ErrorResponse struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Detail  map[string]string `json:"detail"`
}

type Empty struct{}
