package repositories

import (
	"errors"
	"sync"

	"oxo_game/internal/models"
)

var (
	ErrPlayerNotFound = errors.New("player not found")
)

// PlayerRepository is the interface that wraps the basic CRUD operations.
type PlayerRepository interface {
	GetAllPlayers() ([]models.Player, error)
	GetPlayerByID(id int) (*models.Player, error)
	CreatePlayer(player models.Player) (int, error)
	UpdatePlayer(id int, updatedPlayer models.Player) error
	DeletePlayer(id int) error
	DeductBalance(playerID int, amount float64) error
}

// InMemoryPlayerRepository is an example of a repository using in-memory storage.
type InMemoryPlayerRepository struct {
	mu      sync.RWMutex
	players map[int]models.Player
	autoID  int
}

// NewInMemoryPlayerRepository creates a new InMemoryPlayerRepository.
func NewInMemoryPlayerRepository() *InMemoryPlayerRepository {
	return &InMemoryPlayerRepository{
		players: make(map[int]models.Player),
		autoID:  0,
	}
}

// GetAllPlayers returns all players.
func (r *InMemoryPlayerRepository) GetAllPlayers() ([]models.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	players := make([]models.Player, 0, len(r.players))
	for _, player := range r.players {
		players = append(players, player)
	}
	return players, nil
}

// GetPlayerByID returns the player with the given ID.
func (r *InMemoryPlayerRepository) GetPlayerByID(id int) (*models.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	player, ok := r.players[id]
	if !ok {
		return nil, ErrPlayerNotFound
	}
	return &player, nil
}

// CreatePlayer adds a new player and returns the new player's ID.
func (r *InMemoryPlayerRepository) CreatePlayer(player models.Player) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.autoID++
	player.ID = r.autoID
	r.players[player.ID] = player
	return player.ID, nil
}

// UpdatePlayer updates the player with the given ID.
func (r *InMemoryPlayerRepository) UpdatePlayer(id int, updatedPlayer models.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.players[id]; !ok {
		return ErrPlayerNotFound
	}
	updatedPlayer.ID = id
	r.players[id] = updatedPlayer
	return nil
}

// DeletePlayer deletes the player with the given ID.
func (r *InMemoryPlayerRepository) DeletePlayer(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.players[id]; !ok {
		return ErrPlayerNotFound
	}
	delete(r.players, id)
	return nil
}
func (r *InMemoryPlayerRepository) DeductBalance(playerID int, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	player, ok := r.players[playerID]
	if !ok {
		return errors.New("player not found")
	}

	if player.Balance < amount {
		return errors.New("insufficient balance")
	}

	player.Balance -= amount
	return nil
}
