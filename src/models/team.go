package models

type Team struct {
	ID        int    `json:"id"`
	InhouseID int    `json:"inhouse_id"`
	Name      string `json:"name"`
	Players   []Player
}
