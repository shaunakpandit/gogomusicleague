package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Song struct {
	name      string
	roundName string
	score     int
	url       string
}

func SearchSongs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	songName := r.URL.Query().Get("songName")
	songName = songName

	var songs []Song
	rows, err := db.Query(`
		select 
			s.title as title, 
			sum(v.points_assigned) as score,
			r.name as round_name,
			r.playlist_url as playlist_url
		from submissions as s
		join votes as v on s.spotify_uri = v.spotify_uri and s.round_id = v.round_id
		join rounds as r on s.round_id = r.id
		group by title, r.name, r.playlist_url
		`)
	if err != nil {
		http.Error(w, "songs failed", 500)
	}
	defer rows.Close()

	for rows.Next() {
		var song Song
		if err := rows.Scan(
			&song.name,
			&song.score,
			&song.roundName,
			&song.url,
		); err != nil {
			http.Error(w, "songs failed", 500)
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "songs failed", 500)
	}

	json.NewEncoder(w).Encode(songs)
}
