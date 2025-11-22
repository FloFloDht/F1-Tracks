package config

import (
    "os"
)

type Config struct {
    Port        string
    MongoURI    string
    MongoDBName string
}

func Load() *Config {
    return &Config{
        Port:        getEnv("API_PORT", "8080"),
        MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
        MongoDBName: getEnv("MONGO_DB", "f1_circuits"),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}