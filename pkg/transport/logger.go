package transport

import (
	"app/pkg/logger"
	"app/pkg/utils"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type HttpEntity struct {
	HttpStatus int `json:"http_status,omitempty"`

	RequestUri      string      `json:"request_uri,omitempty"`
	RequestBody     string      `json:"request_body,omitempty"`
	RequestHeader   http.Header `json:"request_header,omitempty"`
	RequestTime     string      `json:"request_time,omitempty"`
	RequestMethod   string      `json:"request_method,omitempty"`
	RequestDuration string      `json:"request_duration,omitempty"`

	ResponseBody   string      `json:"response_body,omitempty"`
	ResponseHeader http.Header `json:"response_header,omitempty"`
	ResponseTime   string      `json:"response_time,omitempty"`
	Host           string      `json:"host,omitempty"`

	Errors string `json:"errors,omitempty"`
}

var _ Middleware = &Logger{}

type Logger struct{}

func (t *Logger) Handle(next http.RoundTripper) http.RoundTripper {
	return RoundTripperFunc(func(request *http.Request) (*http.Response, error) {
		ctx := request.Context()
		start := time.Now()

		httpEntity := &HttpEntity{
			RequestUri:    request.URL.RequestURI(),
			Host:          request.URL.Host,
			RequestMethod: request.Method,
			RequestHeader: request.Header,
			RequestTime:   start.Format("2006-01-02 15:04:05"),
		}

		defer func() {
			httpEntity.ResponseTime = time.Now().Format("2006-01-02 15:04:05")
			httpEntity.RequestDuration = time.Since(start).String()
			log.Printf("httpEntity: %+v", httpEntity)
			logger.PrintMap(ctx, toMap(httpEntity))
		}()

		response, err := t.roundTrip(request, next, httpEntity)
		if err != nil {
			httpEntity.Errors = err.Error()
			return nil, err
		}
		return response, nil
	})
}

func (t *Logger) roundTrip(request *http.Request, next http.RoundTripper, httpEntity *HttpEntity) (*http.Response, error) {
	if request.Body != nil {
		// read body wrap with a new reader
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		request.Body = io.NopCloser(bytes.NewBuffer(body))
		httpEntity.RequestBody = string(body)
	}

	response, err := next.RoundTrip(request)
	if err != nil {
		return nil, err
	}

	// read body wrap with a new reader
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body = io.NopCloser(bytes.NewBuffer(body))
	httpEntity.ResponseBody = string(body)
	httpEntity.ResponseHeader = response.Header
	httpEntity.HttpStatus = response.StatusCode
	return response, nil
}

func toMap(entity *HttpEntity) map[string]string {
	return map[string]string{
		"request_uri":      entity.RequestUri,
		"http_status":      strconv.Itoa(entity.HttpStatus),
		"request_body":     entity.RequestBody,
		"request_header":   utils.JsonEncode(entity.RequestHeader),
		"request_time":     entity.RequestTime,
		"request_method":   entity.RequestMethod,
		"request_duration": entity.RequestDuration,
		"response_body":    entity.ResponseBody,
		"response_header":  utils.JsonEncode(entity.ResponseHeader),
		"response_time":    entity.ResponseTime,
		"host":             entity.Host,
		"errors":           entity.Errors,
		"msg":              fmt.Sprintf("%s:%v %s", entity.RequestMethod, entity.HttpStatus, entity.RequestUri),
	}
}
