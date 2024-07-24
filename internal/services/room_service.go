package services

import (
	"errors"
	"sync"

	"oxo_game/internal/models"
	"oxo_game/internal/repositories"
)

type RoomService interface {
	GetAllRooms() ([]models.Room, error)
	GetRoomByID(id int) (*models.Room, error)
	CreateRoom(name, description string) (int, error)
	UpdateRoom(id int, name, description string) error
	DeleteRoom(id int) error
}

type roomService struct {
	roomRepo repositories.RoomRepository
	mu       sync.RWMutex
}

func NewRoomService(repo repositories.RoomRepository) RoomService {
	return &roomService{
		roomRepo: repo,
	}
}

func (s *roomService) GetAllRooms() ([]models.Room, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.roomRepo.GetAllRooms()
}

func (s *roomService) GetRoomByID(id int) (*models.Room, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.roomRepo.GetRoomByID(id)
}

func (s *roomService) CreateRoom(name, description string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if room with the same name already exists
	allRooms, _ := s.roomRepo.GetAllRooms()
	for _, room := range allRooms {
		if room.Name == name {
			return 0, errors.New("room with the same name already exists")
		}
	}

	room := models.Room{
		Name:        name,
		Description: description,
	}

	return s.roomRepo.CreateRoom(room)
}

func (s *roomService) UpdateRoom(id int, name, description string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	room, err := s.roomRepo.GetRoomByID(id)
	if err != nil {
		return err
	}

	room.Name = name
	room.Description = description

	return s.roomRepo.UpdateRoom(id, *room)
}

func (s *roomService) DeleteRoom(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.roomRepo.DeleteRoom(id)
}
