package ginx

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"app/pkg/logger"
	"app/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HttpEntity struct {
	HttpStatus int `json:"http_status,omitempty"`

	RequestUri      string `json:"request_uri,omitempty"`
	RequestBody     string `json:"request_body,omitempty"`
	RequestHeader   string `json:"request_header,omitempty"`
	RequestTime     string `json:"request_time,omitempty"`
	RequestMethod   string `json:"request_method,omitempty"`
	RequestDuration string `json:"request_duration,omitempty"`

	ResponseBody   string `json:"response_body,omitempty"`
	ResponseHeader string `json:"response_header,omitempty"`
	ResponseTime   string `json:"response_time,omitempty"`
	Host           string `json:"host,omitempty"`

	Errors string `json:"errors,omitempty"`
}

type ginWriterWrapper struct {
	gin.ResponseWriter
	buf            *bytes.Buffer
	statusCode     int
	responseHeader http.Header
}

func (g *ginWriterWrapper) Write(b []byte) (int, error) {
	g.buf.Write(b)
	return g.ResponseWriter.Write(b)
}

func (g *ginWriterWrapper) WriteHeader(statusCode int) {
	g.statusCode = statusCode
	g.ResponseWriter.WriteHeader(statusCode)
}

func (g *ginWriterWrapper) Header() http.Header {
	return g.ResponseWriter.Header()
}

var AccessLog = func(ignorePathPrefix []string, desensitize Desensitize) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		for _, prefix := range ignorePathPrefix {
			if strings.HasPrefix(path, prefix) {
				c.Next()
				return
			}
		}
		startTime := time.Now()
		wrapper := &ginWriterWrapper{
			statusCode:     c.Writer.Status(),
			ResponseWriter: c.Writer,
			buf:            &bytes.Buffer{},
		}
		c.Writer = wrapper
		entity := &HttpEntity{
			RequestUri:    c.Request.URL.RequestURI(),
			RequestHeader: utils.JsonEncode(c.Request.Header),
			RequestTime:   startTime.Format("2006-01-02 15:04:05"),
			RequestMethod: c.Request.Method,
			Host:          c.Request.URL.Host,
		}

		requestBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			entity.Errors = err.Error()
		}
		if desensitize != nil {
			entity.RequestBody = string(desensitize.Desensitize(path, requestBody))
		} else {
			entity.RequestBody = string(requestBody)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		c.Next()
		endTime := time.Now()
		entity.ResponseTime = endTime.Format("2006-01-02 15:04:05")
		entity.RequestDuration = endTime.Sub(startTime).String()
		entity.ResponseBody = string(wrapper.buf.Bytes())
		entity.HttpStatus = wrapper.statusCode
		entity.ResponseHeader = utils.JsonEncode(wrapper.Header())

		m := toMap(entity)
		logger.PrintMap(c.Request.Context(), m)
	}
}

func toMap(m *HttpEntity) map[string]string {
	return map[string]string{
		"http_status":      strconv.Itoa(m.HttpStatus),
		"request_uri":      m.RequestUri,
		"request_body":     m.RequestBody,
		"request_header":   m.RequestHeader,
		"request_time":     m.RequestTime,
		"request_method":   m.RequestMethod,
		"request_duration": m.RequestDuration,
		"response_body":    m.ResponseBody,
		"response_header":  m.ResponseHeader,
		"response_time":    m.ResponseTime,
		"host":             m.Host,
		"errors":           m.Errors,
		"msg":              fmt.Sprintf("%s:%v %s", m.RequestMethod, m.HttpStatus, m.RequestUri),
	}
}
