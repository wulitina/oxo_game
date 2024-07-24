package repositories

import (
	"errors"
	"sync"

	"oxo_game/internal/models"
)

var (
	ErrRoomNotFound = errors.New("room not found")
)

// RoomRepository is the interface that wraps the basic CRUD operations for rooms.
type RoomRepository interface {
	GetAllRooms() ([]models.Room, error)
	GetRoomByID(id int) (*models.Room, error)
	CreateRoom(room models.Room) (int, error)
	UpdateRoom(id int, updatedRoom models.Room) error
	DeleteRoom(id int) error
}

// InMemoryRoomRepository is an example of a repository using in-memory storage.
type InMemoryRoomRepository struct {
	mu     sync.RWMutex
	rooms  map[int]models.Room
	autoID int
}

// NewInMemoryRoomRepository creates a new InMemoryRoomRepository.
func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms:  make(map[int]models.Room),
		autoID: 0,
	}
}

// GetAllRooms returns all rooms.
func (r *InMemoryRoomRepository) GetAllRooms() ([]models.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rooms := make([]models.Room, 0, len(r.rooms))
	for _, room := range r.rooms {
		rooms = append(rooms, room)
	}
	return rooms, nil
}

// GetRoomByID returns the room with the given ID.
func (r *InMemoryRoomRepository) GetRoomByID(id int) (*models.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	room, ok := r.rooms[id]
	if !ok {
		return nil, ErrRoomNotFound
	}
	return &room, nil
}

// CreateRoom adds a new room and returns the new room's ID.
func (r *InMemoryRoomRepository) CreateRoom(room models.Room) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.autoID++
	room.ID = r.autoID
	r.rooms[room.ID] = room
	return room.ID, nil
}

// UpdateRoom updates the room with the given ID.
func (r *InMemoryRoomRepository) UpdateRoom(id int, updatedRoom models.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.rooms[id]; !ok {
		return ErrRoomNotFound
	}
	updatedRoom.ID = id
	r.rooms[id] = updatedRoom
	return nil
}

// DeleteRoom deletes the room with the given ID.
func (r *InMemoryRoomRepository) DeleteRoom(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.rooms[id]; !ok {
		return ErrRoomNotFound
	}
	delete(r.rooms, id)
	return nil
}
