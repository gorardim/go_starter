package transport

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var _ http.RoundTripper = &Mock{}

func NewMockClient(suits []HttpSuit) *http.Client {
	return &http.Client{
		Transport: Chain(
			&Mock{
				suits: suits,
			},
			&Debug{},
		),
	}
}

type HttpSuit struct {
	Path         string
	Query        map[string]string
	Status       int
	ResponseBody string
}

type Mock struct {
	suits []HttpSuit
}

func (m *Mock) RoundTrip(request *http.Request) (*http.Response, error) {
	for _, suit := range m.suits {
		if !m.match(request, suit) {
			continue
		}

		status := 200
		if suit.Status != 0 {
			status = suit.Status
		}
		return &http.Response{
			StatusCode: status,
			Body:       io.NopCloser(bytes.NewBufferString(suit.ResponseBody)),
		}, nil
	}
	return nil, fmt.Errorf("not found mock for %s", request.URL.Path)
}

func (m *Mock) match(request *http.Request, suit HttpSuit) bool {
	// path match
	if suit.Path != request.URL.Path {
		return false
	}

	// query match
	if len(suit.Query) > 0 {
		for key, value := range suit.Query {
			if request.URL.Query().Get(key) != value {
				return false
			}
		}
	}
	return true
}
