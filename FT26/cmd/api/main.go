package main

import (
	"log"

	"FT26/internal/config"
	"FT26/internal/mongo"
	"FT26/internal/http"
)

func main() {
    cfg := config.Load()

    mongoClient, err := mongo.NewClient(cfg)
    if err != nil {
        log.Fatalf("MongoDB connection error: %v", err)
    }

    router := http.NewRouter(mongoClient, cfg)

    log.Printf("API listening on port %s", cfg.Port)
    router.Run(":" + cfg.Port)
}
