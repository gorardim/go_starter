package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var _ Middleware = &Debug{}

type Debug struct{}

func (d *Debug) Handle(next http.RoundTripper) http.RoundTripper {
	return RoundTripperFunc(func(request *http.Request) (*http.Response, error) {
		fmt.Println("------------------ request start ------------------")
		fmt.Printf("-> %s: %s\n", request.Method, request.URL.String())
		fmt.Println("-> header:\n", jsonPrettyEncode(request.Header))
		if request.Body != nil {
			body, err := io.ReadAll(request.Body)
			if err != nil {
				return nil, err
			}
			request.Body = io.NopCloser(bytes.NewBuffer(body))
			fmt.Printf("-> body:\n%s\n", string(body))
		}
		fmt.Println("------------------ request end ------------------")
		response, err := next.RoundTrip(request)
		if err != nil {
			return nil, err
		}
		fmt.Println("------------------ response start ------------------")
		fmt.Println("<- status:", response.StatusCode)
		fmt.Printf("<- header:\n%s\n", jsonPrettyEncode(response.Header))
		if response.Body != nil {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			response.Body = io.NopCloser(bytes.NewBuffer(body))
			fmt.Println("<- body:\n", string(body))
		}
		fmt.Println("------------------ response end ------------------")
		return response, nil
	})
}

func jsonPrettyEncode(v interface{}) string {
	indent, _ := json.Marshal(v)
	return string(indent)
}
