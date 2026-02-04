package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"docvault-backend/internal/api/handler"
	"docvault-backend/internal/config"
	"docvault-backend/internal/fs"
	"docvault-backend/internal/logger"
	"docvault-backend/internal/middleware"
	"docvault-backend/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// Initialize logger
	logCfg := logger.Config{
		Level:     cfg.Log.Level,
		Format:    cfg.Log.Format,
		Output:    cfg.Log.Output,
		AddSource: cfg.Log.AddSource,
	}
	if err := logger.Init(logCfg); err != nil {
		logger.Error("Failed to initialize logger", "error", err)
		os.Exit(1)
	}

	logger.Info("Starting DocVault API",
		"address", cfg.GetAddress(),
		"storage", cfg.Storage.RootDir,
		"log_level", cfg.Log.Level,
	)

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Ensure storage directory exists
	if err := os.MkdirAll(cfg.Storage.RootDir, 0755); err != nil {
		logger.Error("Failed to create storage directory", "error", err)
		os.Exit(1)
	}

	// Initialize layers
	docFS := fs.New(cfg.Storage.RootDir)
	docSvc := service.New(docFS)
	docHandler := handler.New(docSvc)

	// Setup router
	r := gin.New() // Use gin.New() to avoid default logger middleware

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
	logger.Info("Server listening",
		"address", addr,
		"openapi", "http://"+addr+"/openapi.yaml",
	)

	// Graceful shutdown
	go func() {
		if err := r.Run(addr); err != nil {
			logger.Error("Server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
}
