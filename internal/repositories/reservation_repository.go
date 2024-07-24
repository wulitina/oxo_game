package repositories

import (
	"errors"
	"sync"
	"time"

	"oxo_game/internal/models"
)

type ReservationRepository interface {
	Create(reservation *models.Reservation) (int, error)
	GetById(id int) (*models.Reservation, error)
	List() []*models.Reservation
	ListByRoomAndDate(roomID int, date time.Time) []*models.Reservation
	Delete(id int) error
}

type InMemoryReservationRepository struct {
	mu           sync.RWMutex
	reservations map[int]*models.Reservation
	autoID       int
}

func NewInMemoryReservationRepository() *InMemoryReservationRepository {
	return &InMemoryReservationRepository{
		reservations: make(map[int]*models.Reservation),
		autoID:       0,
	}
}

func (r *InMemoryReservationRepository) Create(reservation *models.Reservation) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.autoID++
	reservation.ID = r.autoID
	reservation.CreatedAt = time.Now()
	r.reservations[reservation.ID] = reservation
	return reservation.ID, nil
}

func (r *InMemoryReservationRepository) GetById(id int) (*models.Reservation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	reservation, ok := r.reservations[id]
	if !ok {
		return nil, errors.New("reservation not found")
	}
	return reservation, nil
}

func (r *InMemoryReservationRepository) List() []*models.Reservation {
	r.mu.RLock()
	defer r.mu.RUnlock()

	reservations := make([]*models.Reservation, 0, len(r.reservations))
	for _, reservation := range r.reservations {
		reservations = append(reservations, reservation)
	}
	return reservations
}

func (r *InMemoryReservationRepository) ListByRoomAndDate(roomID int, date time.Time) []*models.Reservation {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var reservations []*models.Reservation
	for _, reservation := range r.reservations {
		if reservation.RoomID == roomID && reservation.Date.Equal(date) {
			reservations = append(reservations, reservation)
		}
	}
	return reservations
}
func (r *InMemoryReservationRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.reservations[id]; !ok {
		return errors.New("reservation not found")
	}
	delete(r.reservations, id)
	return nil
}
