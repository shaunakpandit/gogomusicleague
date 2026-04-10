package models

import (
	"database/sql"
	"fmt"
	"time"
)

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

func SubmissionsByCompetitorID(id string, db *sql.DB) ([]Submission, error) {
	// A submissions slice to hold data from returned rows.
	var submissions []Submission

	rows, err := db.Query("SELECT * FROM submissions WHERE submitter_id = ? ORDER BY created DESC", id)
	if err != nil {
		return nil, fmt.Errorf("submissionsByComp %q: %v", id, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var sub Submission
		if err := rows.Scan(
			&sub.ID,
			&sub.SpotifyURI,
			&sub.Title,
			&sub.Album,
			&sub.Artists,
			&sub.SubmitterID,
			&sub.Created,
			&sub.Comment,
			&sub.RoundID,
			&sub.VisibleToVoters,
		); err != nil {
			return nil, fmt.Errorf("submissionsByComp %q: %v", id, err)
		}
		submissions = append(submissions, sub)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("submissionsByComp %q: %v", id, err)
	}
	return submissions, nil
}
