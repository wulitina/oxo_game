package models

type Payment struct {
	ID        int     `json:"id"`
	Method    string  `json:"method"`
	Amount    float64 `json:"amount"`
	Details   string  `json:"details"`
	CreatedAt int64   `json:"created_at"`
}
