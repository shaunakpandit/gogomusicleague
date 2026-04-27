package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jedib0t/go-pretty/v6/table"
	"jssp.io/gogomusicleague/models"
)

func UserSongsHandler(db *sql.DB) {
	cmp := getAndPrintCompetitor(db)

	getAndPrintScoresByCompetitor(db, cmp.ID)
}

func getAndPrintCompetitor(db *sql.DB) models.Competitor {
	cmp, err := models.CompetitorByName("shaunakpandit", db)
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
	fmt.Println(t.Render())
	return comps
}
