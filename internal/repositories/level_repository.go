package repositories

import (
	"errors"
	"sync"

	"oxo_game/internal/models"
)

type LevelRepository interface {
	Create(level *models.Level) (int, error)
	GetById(id int) (*models.Level, error)
	List() []*models.Level
}

type InMemoryLevelRepository struct {
	mu     sync.RWMutex
	levels map[int]*models.Level
	autoID int
}

func NewInMemoryLevelRepository() *InMemoryLevelRepository {
	return &InMemoryLevelRepository{
		levels: make(map[int]*models.Level),
		autoID: 0,
	}
}

func (r *InMemoryLevelRepository) Create(level *models.Level) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.autoID++
	level.ID = r.autoID
	r.levels[level.ID] = level
	return level.ID, nil
}

func (r *InMemoryLevelRepository) GetById(id int) (*models.Level, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	level, ok := r.levels[id]
	if !ok {
		return nil, errors.New("level not found")
	}
	return level, nil
}

func (r *InMemoryLevelRepository) List() []*models.Level {
	r.mu.RLock()
	defer r.mu.RUnlock()

	levels := make([]*models.Level, 0, len(r.levels))
	for _, level := range r.levels {
		levels = append(levels, level)
	}
	return levels
}
