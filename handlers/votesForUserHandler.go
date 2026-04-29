package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jedib0t/go-pretty/v6/table"
	"jssp.io/gogomusicleague/models"
)

type VotesForUser struct {
	Votes   []models.PointsPerCompetitor
	Success bool
}

func VotesForUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name")
	tmpl := template.Must(template.ParseFiles("templates/userSongs.html"))

	if name == "" {
		tmpl.Execute(w, VotesForUser{
			nil,
			true,
		})
		return
	}

	cmp := getAndPrintCompetitor(db, name)
	songs := getAndPrintScoresByCompetitor(db, cmp.ID)

	data := VotesForUser{
		Votes:   songs,
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

	t := table.NewWriter()
	t.AppendHeader(table.Row{"ID", "Name"})
	t.AppendRow(table.Row{cmp.ID, cmp.Name})
	fmt.Println(t.Render())
	return cmp
}

func getAndPrintScoresByCompetitor(db *sql.DB, id string) []models.PointsPerCompetitor {
	comps, err := models.PointsAwardedToCompetitorByCompetitor(id, db)
	if err != nil {
		log.Fatal(err)
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"VoterId", "VoterName", "PointsAwarded"})
	for _, c := range comps {
		t.AppendRow(table.Row{c.VoterId, c.VoterName, c.PointsAwarded})

	}
	// fmt.Println(t.Render())
	return comps
}
