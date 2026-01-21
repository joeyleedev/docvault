package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"docvault-backend/internal/api"
)

// Recovery recovers from panics and returns a 500 error
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID := GetRequestID(c)
				log.Printf("[%s] Panic recovered: %v", requestID, err)

				c.JSON(http.StatusInternalServerError, api.ErrorResponse{
					Code:    api.CodeInternalError,
					Message: api.ErrorMessages[api.CodeInternalError],
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
