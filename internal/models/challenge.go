package models

import "time"

// Challenge represents a game challenge entity.
type Challenge struct {
	ID        int       `json:"id"`
	PlayerID  int       `json:"player_id"`
	CreatedAt time.Time `json:"created_at"`
	Won       bool      `json:"won"`
}

// NewChallenge creates a new Challenge instance with initialized fields.
func NewChallenge(playerID int) *Challenge {
	return &Challenge{
		PlayerID:  playerID,
		CreatedAt: time.Now(),
	}
}
