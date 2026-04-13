package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Vote struct {
	ID             int
	SpotifyURI     string
	VoterID        string
	created        time.Time
	PointsAssigned int
	Comment        *string
	RoundID        string
}

func VotesByCompetitor(voterId string, db *sql.DB) (Competitor, error) {
	var cmp Competitor

	row := db.QueryRow("SELECT * FROM votes WHERE voter_id = ?", voterId)
	if err := row.Scan(&cmp.ID, &cmp.Name); err != nil {
		if err == sql.ErrNoRows {
			return cmp, fmt.Errorf("VotesByCompetitor %s: no such competitor", voterId)
		}
		return cmp, fmt.Errorf("VotesByCompetitor %s: %v", voterId, err)
	}
	return cmp, nil
}
