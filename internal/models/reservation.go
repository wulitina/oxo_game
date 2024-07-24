package models

import "time"

type Reservation struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id"`
	Date      time.Time `json:"date"`
	Time      string    `json:"time"`
	PlayerID  int       `json:"player_id"`
	CreatedAt time.Time `json:"created_at"`
}
