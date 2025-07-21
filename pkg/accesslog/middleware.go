// Package accesslog provides a middleware that records every RESTful API call in a log message.
package accesslog

import (
	"net/http"
	"time"
	"Template/pkg/log"
	"github.com/gin-gonic/gin"
)

// Handler returns a middleware that records an access log message for every HTTP request being processed.
func Handler(logger log.Logger) gin.HandlerFunc{
	return func(c *gin.Context) {
		start := time.Now()
		writer := &bodyLogWriter{
			ResponseWriter: c.Writer,
			statusCode:     http.StatusOK,
		}
		c.Writer = writer


		ctx := c.Request.Context()
		ctx = log.WithRequest(ctx, c.Request)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		duration := time.Since(start)
		status := writer.statusCode

		logger.With(ctx, "duration", duration.Milliseconds(), "status", status).
			Infof("%s %s %s %d %d", c.Request.Method, c.Request.URL.Path, c.Request.Proto, status, writer.size)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	statusCode int
	size       int
}

func (w *bodyLogWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}