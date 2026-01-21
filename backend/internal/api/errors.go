package api

import "net/http"

// Error codes
const (
	// Success
	CodeSuccess = "SUCCESS"

	// Document errors (DOC_xxx)
	CodeDocumentNotFound = "DOC_NOT_FOUND"
	CodeDocumentExists   = "DOC_EXISTS"
	CodeDocumentInvalid  = "DOC_INVALID"

	// Request errors (REQ_xxx)
	CodeInvalidRequest  = "REQ_INVALID"
	CodeMissingParameter = "REQ_MISSING_PARAM"
	CodeInvalidJSON     = "REQ_INVALID_JSON"

	// System errors (SYS_xxx)
	CodeInternalError = "SYS_INTERNAL"
	CodeStorageError  = "SYS_STORAGE"
	CodePathInvalid   = "SYS_PATH_INVALID"

	// Validation errors (VAL_xxx)
	CodeValidationError = "VAL_FAILED"
)

// ErrorMessages maps error codes to human-readable messages
var ErrorMessages = map[string]string{
	CodeSuccess:          "Success",
	CodeDocumentNotFound: "Document not found",
	CodeDocumentExists:   "Document already exists",
	CodeDocumentInvalid:  "Invalid document",
	CodeInvalidRequest:   "Invalid request",
	CodeMissingParameter: "Missing required parameter",
	CodeInvalidJSON:      "Invalid JSON format",
	CodeInternalError:    "Internal server error",
	CodeStorageError:     "Storage operation failed",
	CodePathInvalid:      "Invalid path",
	CodeValidationError:  "Validation failed",
}

// GetHTTPStatus returns HTTP status code for given error code
func GetHTTPStatus(code string) int {
	switch code {
	case CodeSuccess:
		return http.StatusOK
	case CodeDocumentNotFound:
		return http.StatusNotFound
	case CodeDocumentExists:
		return http.StatusConflict
	case CodeDocumentInvalid, CodeInvalidRequest, CodeMissingParameter, CodeInvalidJSON, CodeValidationError:
		return http.StatusBadRequest
	case CodePathInvalid:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
