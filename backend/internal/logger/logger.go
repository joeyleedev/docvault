package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

var (
	// defaultLogger is the global logger instance
	defaultLogger *slog.Logger
	once          sync.Once
)

// Config holds logger configuration
type Config struct {
	Level      string `mapstructure:"level"`      // debug, info, warn, error
	Format     string `mapstructure:"format"`     // json, text
	Output     string `mapstructure:"output"`     // stdout, stderr, or file path
	MaxSize    int    `mapstructure:"max_size"`   // max file size in MB (for rotation)
	MaxBackups int    `mapstructure:"max_backups"` // number of backup files
	MaxAge     int    `mapstructure:"max_age"`     // max age in days
	AddSource  bool   `mapstructure:"add_source"` // add source code position
}

// Init initializes the global logger with the given configuration
func Init(cfg Config) error {
	var initErr error
	once.Do(func() {
		initErr = initLogger(cfg)
	})
	return initErr
}

// initLogger creates and sets the global logger
func initLogger(cfg Config) error {
	// Set default values
	if cfg.Level == "" {
		cfg.Level = "info"
	}
	if cfg.Format == "" {
		cfg.Format = "text"
	}
	if cfg.Output == "" {
		cfg.Output = "stdout"
	}

	// Parse log level
	level, err := parseLevel(cfg.Level)
	if err != nil {
		return fmt.Errorf("invalid log level %q: %w", cfg.Level, err)
	}

	// Get output writer
	writer, err := getWriter(cfg)
	if err != nil {
		return fmt.Errorf("failed to get log writer: %w", err)
	}

	// Create handler options
	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	// Create handler based on format
	var handler slog.Handler
	switch cfg.Format {
	case "json":
		handler = slog.NewJSONHandler(writer, opts)
	case "text", "":
		handler = slog.NewTextHandler(writer, opts)
	default:
		return fmt.Errorf("unsupported log format %q (use 'json' or 'text')", cfg.Format)
	}

	// Set default logger
	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)

	return nil
}

// parseLevel converts string level to slog.Level
func parseLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unknown level: %s", level)
	}
}

// getWriter returns the appropriate writer based on config
func getWriter(cfg Config) (io.Writer, error) {
	switch cfg.Output {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	default:
		// Output is a file path
		if err := os.MkdirAll(filepath.Dir(cfg.Output), 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
		file, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		return file, nil
	}
}

// Get returns the default logger instance
func Get() *slog.Logger {
	if defaultLogger == nil {
		// Initialize with defaults if not initialized
		_ = Init(Config{})
	}
	return defaultLogger
}

// With returns a logger with additional key-value pairs
func With(args ...any) *slog.Logger {
	return Get().With(args...)
}

// Debug logs a debug message
func Debug(msg string, args ...any) {
	Get().Debug(msg, args...)
}

// Info logs an info message
func Info(msg string, args ...any) {
	Get().Info(msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...any) {
	Get().Warn(msg, args...)
}

// Error logs an error message
func Error(msg string, args ...any) {
	Get().Error(msg, args...)
}
