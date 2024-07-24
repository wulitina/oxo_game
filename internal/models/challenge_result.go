package models

type ChallengeResult struct {
	ID         int   `json:"id"`
	PlayerID   int   `json:"player_id"`
	WonJackpot bool  `json:"won_jackpot"`
	CreatedAt  int64 `json:"created_at"`
}
