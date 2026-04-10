package models

import "time"

type Vote struct {
	ID             int
	SpotifyURI     string
	VoterID        string
	created        time.Time
	PointsAssigned int
	Comment        *string
	RoundID        string
}
