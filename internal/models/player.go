package models

type Player struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Level   *Level  `json:"level"`
	Balance float64 `json:"balance"`
}
