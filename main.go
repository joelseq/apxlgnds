package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joelseq/apxlgnds/internal/cache"
	"github.com/joelseq/apxlgnds/internal/calendar"
	"github.com/joelseq/apxlgnds/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		log.Fatalln("env var GOOGLE_API_KEY not found")
	}

	cache := cache.NewCache(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PASSWORD"))
	calendarService := calendar.NewService(apiKey)
	port := getEnvWithDefault("PORT", "8080")
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	srv := server.NewServer(port, cache, calendarService, logger)

	srv.Start()
}

func getEnvWithDefault(env, defaultValue string) string {
	if value, exists := os.LookupEnv(env); exists {
		return value
	}
	return defaultValue
}
