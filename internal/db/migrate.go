package db

import (
    "io/ioutil"
    "log"
)

func Migrate() {
    files := []string{
        "migrations/001_create_tables.sql",
        "migrations/002_seed_initial_data.sql",
    }

    for _, file := range files {
        content, err := ioutil.ReadFile(file)
        if err != nil {
            log.Fatalf("Impossible de lire %s : %v", file, err)
        }

        if _, err := DB.Exec(string(content)); err != nil {
            log.Fatalf("Erreur migration %s : %v", file, err)
        }
    }

    log.Println("Migrations SQLite exécutées ✔️")
}