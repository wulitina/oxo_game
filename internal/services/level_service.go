package services

import (
	"errors"
	"oxo_game/internal/models"
	"oxo_game/internal/repositories"
	"sync"
)

type LevelService interface {
	CreateLevel(name string) (int, error)
	GetLevelByID(id int) (*models.Level, error)
	GetAllLevels() ([]*models.Level, error)
}

type levelService struct {
	levelRepo repositories.LevelRepository
	mu        sync.RWMutex
}

func NewLevelService(repo repositories.LevelRepository) LevelService {
	return &levelService{
		levelRepo: repo,
	}
}

func (s *levelService) CreateLevel(name string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if level with the same name already exists
	allLevels := s.levelRepo.List()
	for _, level := range allLevels {
		if level.Name == name {
			return 0, errors.New("level with the same name already exists")
		}
	}

	// Create new level
	level := &models.Level{Name: name}
	return s.levelRepo.Create(level)
}

func (s *levelService) GetLevelByID(id int) (*models.Level, error) {
	return s.levelRepo.GetById(id)
}

func (s *levelService) GetAllLevels() ([]*models.Level, error) {
	return s.levelRepo.List(), nil
}
