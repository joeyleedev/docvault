package api

import "time"

// CreateDocumentRequest represents a request to create a document
type CreateDocumentRequest struct {
	Name    string `json:"name" binding:"required" example:"my-document"`
	Content string `json:"content" example:"# My Document\n\nThis is the content."`
}

// UpdateDocumentRequest represents a request to update a document
type UpdateDocumentRequest struct {
	Content string `json:"content" binding:"required" example:"# Updated Content\n\nNew content here."`
}

// Document represents a document with its content
type Document struct {
	Name      string `json:"name" example:"my-document"`
	Content   string `json:"content" example:"# My Document\n\nContent here."`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Size      int64  `json:"size,omitempty" example:"1024"`
}

// DocumentMeta represents document metadata without content
type DocumentMeta struct {
	Name      string    `json:"name" example:"my-document.md"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-15T10:30:00Z"`
	Size      int64     `json:"size" example:"1024"`
}

// DocumentListResponse represents a list of documents
type DocumentListResponse struct {
	Total int            `json:"total" example:"10"`
	Items []DocumentMeta `json:"items"`
}
