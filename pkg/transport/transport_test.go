package transport

import (
	"context"
	"io"
	"net/http"
	"testing"

	"app/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestX(t *testing.T) {

	client := http.Client{}
	client.Transport = Chain(
		http.DefaultTransport,
		&Logger{},
		&Debug{},
	)

	req, err := http.NewRequestWithContext(logger.NewLoggerContext(context.Background()), "GET", "https://api.pay.verystar.net/ping.json", nil)
	assert.NoError(t, err)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// read body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
}
