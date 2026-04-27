package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jedib0t/go-pretty/v6/table"
	"jssp.io/gogomusicleague/models"
)

type UserSongsData struct {
	Songs   []models.PointsPerCompetitor
	Success bool
}

func UserSongsHandler(name string, db *sql.DB) UserSongsData {
	cmp := getAndPrintCompetitor(db, name)
	songs := getAndPrintScoresByCompetitor(db, cmp.ID)

	return UserSongsData{
		Songs:   songs,
		Success: true,
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
