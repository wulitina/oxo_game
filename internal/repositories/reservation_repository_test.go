package repositories

import (
	"testing"
	"time"

	"oxo_game/internal/models"
)

func TestInMemoryReservationRepository_CRUDOperations(t *testing.T) {
	repo := NewInMemoryReservationRepository()

	// Create a reservation
	reservation := &models.Reservation{
		RoomID:   1,
		Date:     time.Date(2024, 7, 5, 0, 0, 0, 0, time.UTC),
		Time:     "10:00 AM",
		PlayerID: 1,
	}

	id, err := repo.Create(reservation)
	if err != nil {
		t.Fatalf("Error creating reservation: %v", err)
	}

	// Set the expected ID and CreatedAt
	reservation.ID = id
	reservation.CreatedAt = time.Now()

	// Get reservation by ID
	createdReservation, err := repo.GetById(id)
	if err != nil {
		t.Fatalf("Error fetching reservation by ID: %v", err)
	}

	// Check if the created reservation matches the expected reservation
	if !reservationsAreEqual(reservation, createdReservation) {
		t.Errorf("Created reservation does not match expected. Expected %+v, got %+v", reservation, createdReservation)
	}

	// List all reservations
	allReservations := repo.List()
	if len(allReservations) != 1 {
		t.Fatalf("Expected 1 reservation, got %d", len(allReservations))
	}
	if !reservationsAreEqual(reservation, allReservations[0]) {
		t.Errorf("Listed reservation does not match expected. Expected %+v, got %+v", reservation, allReservations[0])
	}

	// List reservations by room and date
	roomDateReservations := repo.ListByRoomAndDate(reservation.RoomID, reservation.Date)
	if len(roomDateReservations) != 1 {
		t.Fatalf("Expected 1 reservation, got %d", len(roomDateReservations))
	}
	if !reservationsAreEqual(reservation, roomDateReservations[0]) {
		t.Errorf("Room/date reservation does not match expected. Expected %+v, got %+v", reservation, roomDateReservations[0])
	}

	// Delete the reservation
	err = repo.Delete(reservation.ID)
	if err != nil {
		t.Fatalf("Error deleting reservation: %v", err)
	}

	// Verify reservation deletion by trying to fetch it again
	_, err = repo.GetById(id)
	if err == nil {
		t.Errorf("Expected reservation to be deleted, but it still exists")
	} else if err.Error() != "reservation not found" {
		t.Errorf("Expected 'reservation not found' error, got %v", err)
	}
}

// reservationsAreEqual checks if two reservations are equal considering their fields.
func reservationsAreEqual(r1, r2 *models.Reservation) bool {
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.ID == r2.ID &&
		r1.RoomID == r2.RoomID &&
		r1.Date.Equal(r2.Date) &&
		r1.Time == r2.Time &&
		r1.PlayerID == r2.PlayerID &&
		r1.CreatedAt.Equal(r2.CreatedAt)
}
