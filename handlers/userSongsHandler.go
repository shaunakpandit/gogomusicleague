package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"jssp.io/gogomusicleague/models"
)

type UserSongsData struct {
	Songs   []models.Submission
	Success bool
}

func UserSongsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name")
	tmpl := template.Must(template.ParseFiles("templates/userSongs.html"))

	if name == "" {
		tmpl.Execute(w, UserSongsData{
			nil,
			true,
		})
		return
	}

	cmp := getCompetitor(db, name)
	songs := getSongsForUser(db, cmp.ID)

	data := UserSongsData{
		Songs:   songs,
		Success: true,
	}

	tmplErr := tmpl.Execute(w, data)
	if tmplErr != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func getCompetitor(db *sql.DB, name string) models.Competitor {
	cmp, err := models.CompetitorByName(name, db)
	if err != nil {
		log.Fatal(err)
	}

	return cmp
}

func getSongsForUser(db *sql.DB, id string) []models.Submission {
	subs, err := models.SubmissionsByCompetitorID(id, db)
	if err != nil {
		log.Fatal(err)
	}

	return subs
}
