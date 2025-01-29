package transport

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMock_Handle(t *testing.T) {
	t.Run("test path", func(t *testing.T) {
		client := NewMockClient([]HttpSuit{
			{
				Path:         "/",
				ResponseBody: "hello",
			},
		})
		req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
		assert.NoError(t, err)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, "hello", string(body))
	})

	t.Run("test path and query", func(t *testing.T) {
		client := NewMockClient([]HttpSuit{
			{
				Path: "/",
				Query: map[string]string{
					"a": "1",
				},
				ResponseBody: "hello",
			},
		})
		req, err := http.NewRequest("GET", "http://localhost:8080/?a=1", nil)
		assert.NoError(t, err)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, "hello", string(body))
	})
}
