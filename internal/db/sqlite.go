package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitSqlite(path string) {
	var err error

	DB, err = sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("Erreur SQLite: %v", err)
	}

	if _, err := DB.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		log.Fatalf("Erreur PRAGMA: %v", err)
	}
}
