package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

const RequestIDKey = "X-Request-ID"

// RequestID adds a unique request ID to each request
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDKey)
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Set(RequestIDKey, requestID)
		c.Writer.Header().Set(RequestIDKey, requestID)
		c.Next()
	}
}

// GetRequestID returns the request ID from the context
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return ""
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}
