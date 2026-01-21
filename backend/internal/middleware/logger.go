package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger logs HTTP requests
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

		log.Printf("[%s] %s %s %s %d %v",
			requestID,
			method,
			path,
			query,
			status,
			latency,
		)
	}
}
