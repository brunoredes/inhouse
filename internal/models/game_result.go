package models

import "time"

type GameResult struct {
	ID          int       `json:"id"`
	InhouseID   int       `json:"inhouse_id"`
	WinningTeam string    `json:"winning_team"`
	CompletedAt time.Time `json:"completed_at"`
}
