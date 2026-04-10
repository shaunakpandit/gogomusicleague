package models

import "time"

type Round struct {
	ID          string
	Created     time.Time
	Name        string
	Description *string
	PlaylistURL *string
}
