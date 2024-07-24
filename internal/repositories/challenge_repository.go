package repositories

import (
	"errors"
	"sync"
	"time"

	"oxo_game/internal/models"
)

var (
	ErrChallengeNotFound = errors.New("challenge not found")
)

type ChallengeRepository interface {
	Create(challenge *models.Challenge) (int, error)
	GetById(id int) (*models.Challenge, error)
	ListByPlayer(playerID int) []*models.Challenge
	ListLatest(n int) []*models.Challenge
}

type InMemoryChallengeRepository struct {
	mu         sync.RWMutex
	challenges map[int]*models.Challenge
	autoID     int
}

func NewInMemoryChallengeRepository() *InMemoryChallengeRepository {
	return &InMemoryChallengeRepository{
		challenges: make(map[int]*models.Challenge),
		autoID:     0,
	}
}

func (r *InMemoryChallengeRepository) Create(challenge *models.Challenge) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.autoID++
	challenge.ID = r.autoID
	challenge.CreatedAt = time.Now()
	r.challenges[challenge.ID] = challenge
	return challenge.ID, nil
}

func (r *InMemoryChallengeRepository) GetById(id int) (*models.Challenge, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	challenge, ok := r.challenges[id]
	if !ok {
		return nil, errors.New("challenge not found")
	}
	return challenge, nil
}

func (r *InMemoryChallengeRepository) ListByPlayer(playerID int) []*models.Challenge {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var challenges []*models.Challenge
	for _, challenge := range r.challenges {
		if challenge.PlayerID == playerID {
			challenges = append(challenges, challenge)
		}
	}
	return challenges
}

func (r *InMemoryChallengeRepository) ListLatest(n int) []*models.Challenge {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var challenges []*models.Challenge
	count := 0
	for id := len(r.challenges); id > 0; id-- {
		if challenge, ok := r.challenges[id]; ok {
			challenges = append(challenges, challenge)
			count++
			if count == n {
				break
			}
		}
	}
	return challenges
}
