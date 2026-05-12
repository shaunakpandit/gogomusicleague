package handlers

import (
	"database/sql"
	"net/http"
)

type Song struct {
	name      string
	roundName string
	score     int
}

func SearchSongs(w http.ResponseWriter, r *http.Request, db *sql.DB) []Song {
	songName := r.URL.Query().Get("songName")

	query, err := db.Query(`

		`)
}
