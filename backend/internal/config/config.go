package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	Storage StorageConfig `mapstructure:"storage"`
	API     APIConfig     `mapstructure:"api"`
	Log     LogConfig     `mapstructure:"log"`
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// StorageConfig holds document storage configuration
type StorageConfig struct {
	RootDir string `mapstructure:"root_dir"`
}

// APIConfig holds API-specific configuration
type APIConfig struct {
	BasePath string `mapstructure:"base_path"`
	Version  string `mapstructure:"version"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level     string `mapstructure:"level"`
	Format    string `mapstructure:"format"`
	Output    string `mapstructure:"output"`
	AddSource bool   `mapstructure:"add_source"`
}

var cfg *Config

// Load reads configuration from file and environment variables
func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("server.host", "")
	v.SetDefault("server.port", 8088)
	v.SetDefault("storage.root_dir", "./data/md")
	v.SetDefault("api.base_path", "/api")
	v.SetDefault("api.version", "1.0.0")
	v.SetDefault("log.level", "info")
	v.SetDefault("log.format", "text")
	v.SetDefault("log.output", "stdout")
	v.SetDefault("log.add_source", false)

	// Read config file if provided
	if configPath != "" {
		v.SetConfigFile(configPath)
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	} else {
		// Try to find config in default locations
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./configs")
		v.AddConfigPath("/etc/docvault")

		// Read config file if it exists, but don't fail if it doesn't
		_ = v.ReadInConfig()
	}

	// Override with environment variables
	v.SetEnvPrefix("DOCVAULT")
	v.AutomaticEnv()

	// Parse config
	cfg = &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Expand ~ in root_dir
	cfg.Storage.RootDir = expandPath(cfg.Storage.RootDir)

	return cfg, nil
}

// Get returns the loaded configuration
func Get() *Config {
	return cfg
}

// GetAddress returns the server address
func (c *Config) GetAddress() string {
	if c.Server.Host != "" {
		return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
	}
	return fmt.Sprintf(":%d", c.Server.Port)
}

// expandPath expands ~ to home directory
func expandPath(path string) string {
	if len(path) > 0 && path[0] == '~' {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[1:])
	}
	return path
}
