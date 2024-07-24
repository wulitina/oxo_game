package services

import (
	"errors"
	"oxo_game/internal/models"
	"oxo_game/internal/repositories"
)

var (
	ErrPlayerNotFound = errors.New("player not found")
)

type PlayerService struct {
	repo repositories.PlayerRepository
}

func NewPlayerService(repo repositories.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) GetAllPlayers() ([]models.Player, error) {
	return s.repo.GetAllPlayers()
}

func (s *PlayerService) GetPlayerByID(id int) (*models.Player, error) {
	return s.repo.GetPlayerByID(id)
}

func (s *PlayerService) CreatePlayer(player models.Player) (int, error) {
	return s.repo.CreatePlayer(player)
}

func (s *PlayerService) UpdatePlayer(id int, player models.Player) error {
	return s.repo.UpdatePlayer(id, player)
}

func (s *PlayerService) DeletePlayer(id int) error {
	return s.repo.DeletePlayer(id)
}
