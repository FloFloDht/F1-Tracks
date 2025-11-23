package main

import (
	"F1-Tracks/internal/db"
	"F1-Tracks/internal/tracks"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitSqlite("data/f1.db")
	db.Migrate()

	r := gin.Default()

	tracks.RegisterRoutes(r)

	log.Println("API F1 lancée sur :8080")
	r.Run(":8080")
}
