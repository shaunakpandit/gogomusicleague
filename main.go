package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"jssp.io/gogomusicleague/models"

	"github.com/go-sql-driver/mysql"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

func initDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "192.168.4.28:33061"
	cfg.DBName = "musicleague"
	cfg.ParseTime = true

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
