package main

import (
    "FT26/internal/db"
    "FT26/internal/tracks"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    db.InitSqlite("data/f1.db")
    db.Migrate()

    r := gin.Default()

    tracks.RegisterRoutes(r)

    log.Println("API F1 lancée sur :8080")
    r.Run(":8080")
}