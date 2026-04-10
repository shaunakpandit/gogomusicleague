package models

import (
	"database/sql"
	"fmt"
)

type Competitor struct {
	ID   string
	Name string
}

// CompetitorByName queries for the competitor with the specified name
func CompetitorByName(name string, db *sql.DB) (Competitor, error) {
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

// CompetitorByID queries for the competitor with the specified name
func CompetitorByID(id string, db *sql.DB) (Competitor, error) {
	var cmp Competitor

	row := db.QueryRow("SELECT * FROM competitors WHERE id = ?", id)
	if err := row.Scan(&cmp.ID, &cmp.Name); err != nil {
		if err == sql.ErrNoRows {
			return cmp, fmt.Errorf("competitorById %s: no such competitor", id)
		}
		return cmp, fmt.Errorf("competitorById %s: %v", id, err)
	}
	return cmp, nil
}
