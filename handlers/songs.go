package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
)

type Song struct {
	name      string
	roundName string
	score     int
	url       string
}

type SongResponse struct {
	Songs   []Song
	Success bool
}

func SearchSongs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	songName := r.URL.Query().Get("songName")
	songName = songName

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

	var songs []Song
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

	songResponse := SongResponse{
		Songs:   songs,
		Success: true,
	}

	tmpl := template.Must(template.ParseFiles("templates/songs.html"))
	tmplErr := tmpl.Execute(w, songResponse)
	if tmplErr != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

}
