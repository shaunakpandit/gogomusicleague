package competitors

import (
	"database/sql"
	"fmt"
)

type Competitor struct {
	ID   string
	Name string
}

// competitorByName queries for the competitor with the specified name
func competitorByName(name string, db *sql.DB) (Competitor, error) {
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
