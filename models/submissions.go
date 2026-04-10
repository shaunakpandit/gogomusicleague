package models

import "time"

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
