package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"oxo_game/internal/services"
)

type ChallengeHandler struct {
	challengeService services.ChallengeService
}

func NewChallengeHandler(challengeService services.ChallengeService) *ChallengeHandler {
	return &ChallengeHandler{
		challengeService: challengeService,
	}
}

func (h *ChallengeHandler) ParticipateChallenge(c *gin.Context) {
	playerID, err := strconv.Atoi(c.Param("player_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	wonJackpot, err := h.challengeService.ParticipateChallenge(playerID)
	if err != nil {
		if err.Error() == "player is on cooldown" {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"won_jackpot": wonJackpot})
}

func (h *ChallengeHandler) ListLatestChallenges(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil || n <= 0 {
		n = 10 // Default to 10 if n is invalid or not provided
	}

	challenges := h.challengeService.ListLatestChallenges(n)
	c.JSON(http.StatusOK, challenges)
}
