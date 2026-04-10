package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Competitor struct {
	ID   string
	Name string
}

type Round struct {
	ID          string
	Created     time.Time
	Name        string
	Description *string
	PlaylistURL *string
}

type Submission struct {
	ID              int
	SpotifyURI      string
	Title           string
	Album           *string
	Artists         string
	SubmitterID     string
	Created         *time.Time
	Comment         *string
	RoundID         string
	VisibleToVoters bool
}

type Vote struct {
	ID             int
	SpotifyURI     string
	VoterID        string
	created        time.Time
	PointsAssigned int
	Comment        *string
	RoundID        string
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("No .env file found")
	}
	initDb()

	cmp, err := competitorByName("shaunakpandit")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Competitor found: %v\n", cmp)
}

func initDb() {
	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "192.168.4.28:33061"
	cfg.DBName = "musicleague"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// competitorByName queries for the competitor with the specified name
func competitorByName(name string) (Competitor, error) {
	var cmp Competitor

	row := db.QueryRow("SELECT * FROM competitors WHERE name = ?", name)
	if err := row.Scan(&cmp.ID, &cmp.Name); err != nil {
		if err == sql.ErrNoRows {
			return cmp, fmt.Errorf("competitorByName %s: no such competitor", name)
		}
		return cmp, fmt.Errorf("competitorByName %s: %v", name, err)
	}
	return cmp, nil
}
