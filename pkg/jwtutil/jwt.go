package jwtutil

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims[T any] struct {
	Custom *T `json:"cus"`
	jwt.RegisteredClaims
}

func Sign[T any](secret string, duration time.Duration, data *T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims[T]{
		Custom: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	})
	return token.SignedString([]byte(secret))
}

func Parse[T any](secret string, signed string) (*Claims[T], error) {
	c := &Claims[T]{}
	_, err := jwt.ParseWithClaims(signed, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}
