package provider

import (
	"net/http"

	"app/pkg/transport"
	"app/services/internal/config"
)

func NewClient(conf *config.Config) *http.Client {
	return transport.NewClient(
		&transport.Debug{},
		&transport.Logger{},
	)
}
