package config

import (
	"os"
	"strings"
)

type Config struct {
	Port             string
	DatabaseURL      string
	JWTSecret        string
	SupabaseURL      string
	SupabaseKey      string
	FirebaseProjectID string
	AllowedOrigins   []string
}

func Load() *Config {
	allowedOrigins := []string{"http://localhost:5173", "http://localhost:3000"}
	if env := os.Getenv("ALLOWED_ORIGINS"); env != "" {
		allowedOrigins = strings.Split(env, ",")
	}

	return &Config{
		Port:              getEnv("PORT", "8080"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		SupabaseURL:       os.Getenv("SUPABASE_URL"),
		SupabaseKey:       os.Getenv("SUPABASE_KEY"),
		FirebaseProjectID: os.Getenv("FIREBASE_PROJECT_ID"),
		AllowedOrigins:    allowedOrigins,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
