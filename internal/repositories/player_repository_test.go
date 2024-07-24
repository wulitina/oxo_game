package repositories

import (
	"errors"
	"testing"

	"oxo_game/internal/models"
)

func TestInMemoryPlayerRepository_CRUDOperations(t *testing.T) {
	repo := NewInMemoryPlayerRepository()

	// Create a player
	player := models.Player{
		Name:    "Alice",
		Level:   &models.Level{Name: "Beginner"},
		Balance: 100.0,
	}

	id, err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("Error creating player: %v", err)
	}

	// Set the expected ID
	player.ID = id

	// Get player by ID
	createdPlayer, err := repo.GetPlayerByID(id)
	if err != nil {
		t.Fatalf("Error fetching player by ID: %v", err)
	}

	// Check if the created player matches the expected player
	if !playersAreEqual(&player, createdPlayer) {
		t.Errorf("Created player does not match expected. Expected %+v, got %+v", player, *createdPlayer)
	}

	// Update player
	updatedPlayer := *createdPlayer
	updatedPlayer.Name = "Updated Alice"
	updatedPlayer.Balance = 150.0

	err = repo.UpdatePlayer(id, updatedPlayer)
	if err != nil {
		t.Fatalf("Error updating player: %v", err)
	}

	// Get updated player by ID
	updatedPlayerResult, err := repo.GetPlayerByID(id)
	if err != nil {
		t.Fatalf("Error fetching updated player by ID: %v", err)
	}

	// Check if the updated player matches the expected updated player
	if !playersAreEqual(&updatedPlayer, updatedPlayerResult) {
		t.Errorf("Updated player does not match expected. Expected %+v, got %+v", updatedPlayer, *updatedPlayerResult)
	}

	// Delete player
	err = repo.DeletePlayer(id)
	if err != nil {
		t.Fatalf("Error deleting player: %v", err)
	}

	// Verify player deletion by trying to fetch it again
	_, err = repo.GetPlayerByID(id)
	if err == nil {
		t.Errorf("Expected player to be deleted, but it still exists")
	} else if !errors.Is(err, ErrPlayerNotFound) {
		t.Errorf("Expected ErrPlayerNotFound, got %v", err)
	}
}

func TestInMemoryPlayerRepository_DeductBalance(t *testing.T) {
	repo := NewInMemoryPlayerRepository()

	// Create a player with initial balance
	player := models.Player{
		Name:    "Bob",
		Level:   &models.Level{Name: "Intermediate"},
		Balance: 200.0,
	}

	id, err := repo.CreatePlayer(player)
	if err != nil {
		t.Fatalf("Error creating player: %v", err)
	}

	// Deduct balance from player
	deductAmount := 50.0
	err = repo.DeductBalance(id, deductAmount)
	if err != nil {
		t.Fatalf("Error deducting balance: %v", err)
	}

	// Get player after balance deduction
	updatedPlayer, err := repo.GetPlayerByID(id)
	if err != nil {
		t.Fatalf("Error fetching player by ID after balance deduction: %v", err)
	}

	// Check if the player's balance is correctly deducted
	expectedBalance := player.Balance - deductAmount
	if updatedPlayer.Balance != expectedBalance {
		t.Errorf("Expected balance after deduction to be %.2f, but got %.2f", expectedBalance, updatedPlayer.Balance)
	}

	// Attempt to deduct more than the available balance
	insufficientAmount := updatedPlayer.Balance + 10.0
	err = repo.DeductBalance(id, insufficientAmount)
	if err == nil {
		t.Errorf("Expected error for insufficient balance, but got nil")
	} else if err.Error() != "insufficient balance" {
		t.Errorf("Expected 'insufficient balance' error, but got %v", err)
	}

	// Delete player
	err = repo.DeletePlayer(id)
	if err != nil {
		t.Fatalf("Error deleting player: %v", err)
	}
}

// playersAreEqual checks if two players are equal considering their fields, including Level pointer.
func playersAreEqual(p1, p2 *models.Player) bool {
	if p1 == nil || p2 == nil {
		return false
	}
	return p1.ID == p2.ID &&
		p1.Name == p2.Name &&
		levelsAreEqual(p1.Level, p2.Level) &&
		p1.Balance == p2.Balance
}

// levelsAreEqual checks if two levels are equal.
func levelsAreEqual(l1, l2 *models.Level) bool {
	if l1 == nil || l2 == nil {
		return false
	}
	return l1.Name == l2.Name
}
