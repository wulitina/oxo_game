package services

import (
	"time"

	"oxo_game/internal/models"
	"oxo_game/internal/repositories"
)

type ReservationService interface {
	CreateReservation(roomID int, date time.Time, timeSlot string, playerID int) (int, error)
	GetReservationByID(id int) (*models.Reservation, error)
	ListReservations() []*models.Reservation
	ListReservationsByRoomAndDate(roomID int, date time.Time) []*models.Reservation
}

type reservationService struct {
	reservationRepo repositories.ReservationRepository
}

func NewReservationService(repo repositories.ReservationRepository) ReservationService {
	return &reservationService{
		reservationRepo: repo,
	}
}

func (s *reservationService) CreateReservation(roomID int, date time.Time, timeSlot string, playerID int) (int, error) {
	reservation := &models.Reservation{
		RoomID:    roomID,
		Date:      date,
		Time:      timeSlot,
		PlayerID:  playerID,
		CreatedAt: time.Now(),
	}

	id, err := s.reservationRepo.Create(reservation)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *reservationService) GetReservationByID(id int) (*models.Reservation, error) {
	return s.reservationRepo.GetById(id)
}

func (s *reservationService) ListReservations() []*models.Reservation {
	return s.reservationRepo.List()
}

func (s *reservationService) ListReservationsByRoomAndDate(roomID int, date time.Time) []*models.Reservation {
	return s.reservationRepo.ListByRoomAndDate(roomID, date)
}
