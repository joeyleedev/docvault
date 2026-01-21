package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"docvault-backend/internal/api/handler"
	"docvault-backend/internal/config"
	"docvault-backend/internal/fs"
	"docvault-backend/internal/middleware"
	"docvault-backend/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Ensure storage directory exists
	if err := os.MkdirAll(cfg.Storage.RootDir, 0755); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}

	// Initialize layers
	docFS := fs.New(cfg.Storage.RootDir)
	docSvc := service.New(docFS)
	docHandler := handler.New(docSvc)

	// Setup router
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORS())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "DocVault API",
			"version": cfg.API.Version,
		})
	})

	// OpenAPI spec endpoint (serves the YAML spec)
	r.GET("/openapi.yaml", func(c *gin.Context) {
		c.File("./api/v1/openapi.yaml")
	})

	// API routes
	api := r.Group(cfg.API.BasePath)
	docHandler.Register(api)

	// Start server
	addr := cfg.GetAddress()
	log.Printf("Starting DocVault API on %s", addr)
	log.Printf("Storage directory: %s", cfg.Storage.RootDir)
	log.Printf("OpenAPI spec: http://%s/openapi.yaml", addr)

	// Graceful shutdown
	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
