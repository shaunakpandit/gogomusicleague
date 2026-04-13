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

type PointsPerCompetitor struct {
	VoterId       string
	VoterName     string
	PointsAwarded int
}

func PointsAwardedToCompetitorByCompetitor(id string, db *sql.DB) ([]PointsPerCompetitor, error) {
	rows, err := db.Query(`
		select
			v.voter_id,
			vc.name as voter_name,
			sum(v.points_assigned) as points
		from competitors as c
		join submissions as s
			on s.submitter_id = c.id
		join votes as v
			on s.spotify_uri = v.spotify_uri
		and s.round_id = v.round_id
		join competitors as vc
			on vc.id = v.voter_id
		where c.id = ?
		group by v.voter_id, vc.name
		order by points desc;
		`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PointsPerCompetitor

	for rows.Next() {
		var result PointsPerCompetitor
		if err := rows.Scan(&result.VoterId, &result.VoterName, &result.PointsAwarded); err != nil {
			return results, err
		}
		results = append(results, result)
	}
	if err = rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}
