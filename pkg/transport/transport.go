package transport

import "net/http"

type Middleware interface {
	Handle(next http.RoundTripper) http.RoundTripper
}

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (r RoundTripperFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return r(request)
}

func Chain(next http.RoundTripper, middlewares ...Middleware) http.RoundTripper {
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i].Handle(next)
	}
	return RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		return next.RoundTrip(req)
	})
}

func NewClient(middlewares ...Middleware) *http.Client {
	return &http.Client{
		Transport: Chain(
			http.DefaultTransport,
			middlewares...,
		),
	}
}
