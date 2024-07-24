package services

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"oxo_game/internal/models"
	"oxo_game/internal/repositories"
)

const (
	challengeDuration     = 30 * time.Second
	challengeCooldownTime = 1 * time.Minute
	jackpotWinProbability = 1 // 1% chance to win jackpot
)

type ChallengeService interface {
	ParticipateChallenge(playerID int) (bool, error)
	ListLatestChallenges(n int) []*models.Challenge
}

type challengeService struct {
	challengeRepo repositories.ChallengeRepository
	playerRepo    repositories.PlayerRepository // Assuming you have a player repository
	mu            sync.Mutex
}

func NewChallengeService(challengeRepo repositories.ChallengeRepository, playerRepo repositories.PlayerRepository) ChallengeService {
	return &challengeService{
		challengeRepo: challengeRepo,
		playerRepo:    playerRepo,
	}
}

func (s *challengeService) ParticipateChallenge(playerID int) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the player is eligible to participate
	lastChallenge, err := s.getLastChallengeForPlayer(playerID)
	if err != nil && err.Error() != repositories.ErrChallengeNotFound.Error() {
		return false, err
	}
	if lastChallenge != nil {
		if time.Since(lastChallenge.CreatedAt) < challengeCooldownTime {
			return false, errors.New("player is on cooldown")
		}
	}

	// Deduct payment from the player (assuming this operation is successful)
	err = s.playerRepo.DeductBalance(playerID, 20.01)
	if err != nil {
		return false, err
	}

	// Simulate the challenge
	wonJackpot := s.simulateChallenge()

	// Create a new challenge record
	challenge := &models.Challenge{
		PlayerID:  playerID,
		CreatedAt: time.Now(),
		Won:       wonJackpot,
	}
	_, err = s.challengeRepo.Create(challenge)
	if err != nil {
		return false, err
	}

	return wonJackpot, nil
}

func (s *challengeService) ListLatestChallenges(n int) []*models.Challenge {
	return s.challengeRepo.ListLatest(n)
}

func (s *challengeService) getLastChallengeForPlayer(playerID int) (*models.Challenge, error) {
	challenges := s.challengeRepo.ListByPlayer(playerID)
	if len(challenges) == 0 {
		return nil, repositories.ErrChallengeNotFound
	}
	return challenges[len(challenges)-1], nil
}

func (s *challengeService) simulateChallenge() bool {
	// Simulate a 1% chance to win the jackpot
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)
	return chance == 0
}
