package tracks

import (
	"F1-Tracks/internal/db"
	"log"
)

func GetAllTracks() ([]Track, error) {
	rows, err := db.DB.Query(`
        SELECT id, name, country, length_km, turns, created_year FROM tracks
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Track

	for rows.Next() {
		var t Track
		if err := rows.Scan(&t.ID, &t.Name, &t.Country, &t.LengthKm, &t.Turns, &t.CreatedYear); err != nil {
			log.Println("Erreur scan track:", err)
			continue
		}
		list = append(list, t)
	}

	return list, nil
}

func GetTrackByID(id string) (*Track, error) {
	row := db.DB.QueryRow(`
        SELECT id, name, country, length_km, turns, created_year FROM tracks WHERE id = ?
    `, id)

	var t Track
	if err := row.Scan(&t.ID, &t.Name, &t.Country, &t.LengthKm, &t.Turns, &t.CreatedYear); err != nil {
		return nil, err
	}

	return &t, nil
}
