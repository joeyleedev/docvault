package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"docvault-backend/internal/logger"
)

// Logger logs HTTP requests using structured logging
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		requestID := GetRequestID(c)
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Determine log level based on status code
		level := slog.LevelInfo
		if status >= 400 && status < 500 {
			level = slog.LevelWarn
		} else if status >= 500 {
			level = slog.LevelError
		}

		logger.Get().Log(c.Request.Context(), level, "HTTP request",
			slog.String("request_id", requestID),
			slog.String("method", method),
			slog.String("path", path),
			slog.String("query", query),
			slog.Int("status", status),
			slog.Duration("latency", latency),
			slog.String("ip", clientIP),
			slog.String("user_agent", userAgent),
		)
	}
}
