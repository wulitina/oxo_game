package models

type Log struct {
	ID        int    `json:"id"`
	PlayerID  int    `json:"player_id"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
