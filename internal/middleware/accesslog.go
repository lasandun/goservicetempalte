package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// AccessLog returns a Gin middleware that logs one JSON record per request.
// "skip" lets you omit chatty routes (e.g., /healthz, /metrics).
func AccessLog(skip map[string]struct{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // process the request

		route := c.FullPath()
		if route == "" {
			route = c.Request.URL.Path
		}
		if _, ok := skip[route]; ok {
			return
		}

		status := c.Writer.Status()
		lat := time.Since(start)

		// Level by status class
		var logFn func(msg string, args ...any)
		switch {
		case status >= 500:
			logFn = slog.Error
		case status >= 400:
			logFn = slog.Warn
		default:
			logFn = slog.Info
		}

		reqID := c.GetString("req_id") // set by requestid middleware

		reqBytes := c.Request.ContentLength
		respBytes := c.Writer.Size()

		logFn("http_request",
			"ts", time.Now().Format(time.RFC3339Nano),
			"method", c.Request.Method,
			"path", route,
			"raw_path", c.Request.URL.RequestURI(),
			"status", status,
			"lat_ms", lat.Milliseconds(),
			"ip", c.ClientIP(),
			"ua", c.Request.UserAgent(),
			"referer", c.Request.Referer(),
			"bytes_in", reqBytes,
			"bytes_out", respBytes,
			"req_id", reqID,
			"errors", c.Errors.String(), // non-empty if c.Error(...) used upstream
		)
	}
}
