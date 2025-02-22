package models

import (
	"time"
)

type Inhouse struct {
	ID        string    `json:"id"` // ULID instead of int
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"` // "waiting", "ongoing", "completed"
	Players   []Player  `json:"players,omitempty"`
}
