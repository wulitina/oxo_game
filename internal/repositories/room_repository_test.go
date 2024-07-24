package repositories

import (
	"testing"

	"oxo_game/internal/models"
)

func TestInMemoryRoomRepository_CRUDOperations(t *testing.T) {
	repo := NewInMemoryRoomRepository()

	// Create a room
	room := models.Room{
		Name:        "Room 1",
		Description: "This is Room 1",
		Status:      "Available",
	}

	id, err := repo.CreateRoom(room)
	if err != nil {
		t.Fatalf("Error creating room: %v", err)
	}

	// Get room by ID
	createdRoom, err := repo.GetRoomByID(id)
	if err != nil {
		t.Fatalf("Error fetching room by ID: %v", err)
	}

	// Check if the created room matches the expected room
	room.ID = id
	if !roomsAreEqual(&room, createdRoom) {
		t.Errorf("Created room does not match expected. Expected %+v, got %+v", room, createdRoom)
	}

	// Update room
	updatedRoom := *createdRoom
	updatedRoom.Name = "Updated Room 1"
	updatedRoom.Status = "Occupied"

	err = repo.UpdateRoom(id, updatedRoom)
	if err != nil {
		t.Fatalf("Error updating room: %v", err)
	}

	// Get updated room by ID
	updatedRoomResult, err := repo.GetRoomByID(id)
	if err != nil {
		t.Fatalf("Error fetching updated room by ID: %v", err)
	}

	// Check if the updated room matches the expected updated room
	if !roomsAreEqual(&updatedRoom, updatedRoomResult) {
		t.Errorf("Updated room does not match expected. Expected %+v, got %+v", updatedRoom, updatedRoomResult)
	}

	// List all rooms
	allRooms, err := repo.GetAllRooms()
	if err != nil {
		t.Fatalf("Error listing all rooms: %v", err)
	}
	if len(allRooms) != 1 {
		t.Fatalf("Expected 1 room, got %d", len(allRooms))
	}
	if !roomsAreEqual(&updatedRoom, &allRooms[0]) {
		t.Errorf("Listed room does not match expected. Expected %+v, got %+v", updatedRoom, allRooms[0])
	}

	// Delete room
	err = repo.DeleteRoom(id)
	if err != nil {
		t.Fatalf("Error deleting room: %v", err)
	}

	// Verify room deletion by trying to fetch it again
	_, err = repo.GetRoomByID(id)
	if err == nil {
		t.Errorf("Expected room to be deleted, but it still exists")
	} else if err != ErrRoomNotFound {
		t.Errorf("Expected ErrRoomNotFound, got %v", err)
	}
}

// roomsAreEqual checks if two rooms are equal considering their fields.
func roomsAreEqual(r1, r2 *models.Room) bool {
	if r1 == nil || r2 == nil {
		return false
	}
	return r1.ID == r2.ID &&
		r1.Name == r2.Name &&
		r1.Description == r2.Description &&
		r1.Status == r2.Status
}
