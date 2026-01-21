package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response structure
type Response struct {
	Code    string `json:"code" example:"SUCCESS"`
	Message string `json:"message" example:"Operation successful"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Code    string `json:"code" example:"DOC_NOT_FOUND"`
	Message string `json:"message" example:"Document not found"`
	Details string `json:"details,omitempty"`
}

// SuccessJSON returns a JSON response with the default success message
func SuccessJSON(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "Success",
		Data:    data,
	})
}

// SuccessJSONWithMessage returns a JSON response with a custom success message
func SuccessJSONWithMessage(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	})
}

// ErrorJSON returns a JSON error response
func ErrorJSON(c *gin.Context, code string) {
	c.JSON(GetHTTPStatus(code), ErrorResponse{
		Code:    code,
		Message: ErrorMessages[code],
	})
}

// ErrorJSONWithMessage returns a JSON error response with a custom message
func ErrorJSONWithMessage(c *gin.Context, code, message string) {
	c.JSON(GetHTTPStatus(code), ErrorResponse{
		Code:    code,
		Message: message,
	})
}

// ErrorJSONWithDetails returns a JSON error response with details
func ErrorJSONWithDetails(c *gin.Context, code, details string) {
	c.JSON(GetHTTPStatus(code), ErrorResponse{
		Code:    code,
		Message: ErrorMessages[code],
		Details: details,
	})
}

// SuccessJSONCreated returns a JSON response with HTTP 201 status
func SuccessJSONCreated(c *gin.Context, message string, data any) {
	c.JSON(http.StatusCreated, Response{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	})
}
