package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"jssp.io/gogomusicleague/models"
)

type VotesForUser struct {
	Votes   []models.PointsPerCompetitor
	Success bool
}

func VotesForUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name")
	tmpl := template.Must(template.ParseFiles("templates/votesForUser.html"))

	if name == "" {
		tmpl.Execute(w, VotesForUser{
			nil,
			true,
		})
		return
	}

	cmp := getAndPrintCompetitor(db, name)
	votes := getAndPrintScoresByCompetitor(db, cmp.ID)

	data := VotesForUser{
		Votes:   votes,
		Success: true,
	}

	tmplErr := tmpl.Execute(w, data)
	if tmplErr != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func getAndPrintCompetitor(db *sql.DB, name string) models.Competitor {
	cmp, err := models.CompetitorByName(name, db)
	if err != nil {
		log.Fatal(err)
	}

	return cmp
}

func getAndPrintScoresByCompetitor(db *sql.DB, id string) []models.PointsPerCompetitor {
	comps, err := models.PointsAwardedToCompetitorByCompetitor(id, db)
	if err != nil {
		log.Fatal(err)
	}

	return comps
}
