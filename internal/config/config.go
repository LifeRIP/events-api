package config

import (
	"os"
)

// Config representa la configuración de la aplicación
type Config struct {
	Port             string
	MongoURI         string
	MongoDatabase    string
	EventsCollection string
	LogLevel         string
}

// NewConfig crea una nueva instancia de configuración
func NewConfig() *Config {
	return &Config{
		Port:             getEnv("PORT", "8080"),
		MongoURI:         getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDatabase:    getEnv("MONGO_DATABASE", "events_db"),
		EventsCollection: getEnv("EVENTS_COLLECTION", "events"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
	}
}

// getEnv obtiene una variable de entorno o devuelve un valor predeterminado
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
