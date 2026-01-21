package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"docvault-backend/internal/api"
	"docvault-backend/internal/service"
)

type DocHandler struct {
	svc *service.Service
}

func New(svc *service.Service) *DocHandler {
	return &DocHandler{svc: svc}
}

// Register registers all document routes
func (h *DocHandler) Register(r *gin.RouterGroup) {
	docs := r.Group("/docs")
	{
		docs.GET("", h.List)
		docs.POST("", h.Create)
		docs.GET("/:name", h.Get)
		docs.PUT("/:name", h.Update)
		docs.DELETE("/:name", h.Delete)
	}
}

// List handles GET /api/docs - returns a list of all documents
func (h *DocHandler) List(c *gin.Context) {
	docs, err := h.svc.ListDocs()
	if err != nil {
		api.ErrorJSON(c, api.CodeInternalError)
		return
	}
	api.SuccessJSON(c, api.DocumentListResponse{
		Total: len(docs),
		Items: docs,
	})
}

// Get handles GET /api/docs/:name - returns document content
func (h *DocHandler) Get(c *gin.Context) {
	name := c.Param("name")
	content, err := h.svc.ReadDoc(name)
	if err != nil {
		api.ErrorJSON(c, api.CodeDocumentNotFound)
		return
	}
	api.SuccessJSON(c, api.Document{
		Name:    name,
		Content: string(content),
	})
}

// Update handles PUT /api/docs/:name - updates document content
func (h *DocHandler) Update(c *gin.Context) {
	name := c.Param("name")
	var req api.UpdateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.ErrorJSON(c, api.CodeInvalidJSON)
		return
	}
	if err := h.svc.SaveDoc(name, []byte(req.Content)); err != nil {
		api.ErrorJSON(c, api.CodeStorageError)
		return
	}
	c.Status(http.StatusNoContent)
}

// Create handles POST /api/docs - creates a new document
func (h *DocHandler) Create(c *gin.Context) {
	var req api.CreateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" {
		api.ErrorJSON(c, api.CodeInvalidRequest)
		return
	}
	if err := h.svc.CreateDoc(req.Name, []byte(req.Content)); err != nil {
		api.ErrorJSONWithDetails(c, api.CodeDocumentExists, err.Error())
		return
	}
	api.SuccessJSONCreated(c, "Document created", api.Document{
		Name: req.Name,
	})
	c.Header("Location", "/api/docs/"+req.Name)
}

// Delete handles DELETE /api/docs/:name - deletes a document
func (h *DocHandler) Delete(c *gin.Context) {
	name := c.Param("name")
	if err := h.svc.DeleteDoc(name); err != nil {
		api.ErrorJSON(c, api.CodeDocumentNotFound)
		return
	}
	c.Status(http.StatusNoContent)
}
