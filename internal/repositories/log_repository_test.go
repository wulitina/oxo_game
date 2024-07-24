package repositories

import (
	"errors"
	"testing"
	"time"

	"oxo_game/internal/models"
)

func TestInMemoryLogRepository_CRUDOperations(t *testing.T) {
	repo := NewInMemoryLogRepository()

	// Create a log
	log := models.Log{
		PlayerID:  1,
		Action:    "Login",
		Timestamp: time.Now().Unix(),
	}

	id, err := repo.CreateLog(log)
	if err != nil {
		t.Fatalf("Error creating log: %v", err)
	}

	// Set the expected ID
	log.ID = id

	// Get log by ID
	createdLog, err := repo.GetLogByID(id)
	if err != nil {
		t.Fatalf("Error fetching log by ID: %v", err)
	}

	// Check if the created log matches the expected log
	if !logsAreEqual(createdLog, &log) {
		t.Errorf("Created log does not match expected. Expected %+v, got %+v", log, createdLog)
	}

	// Get all logs
	allLogs, err := repo.GetAllLogs()
	if err != nil {
		t.Fatalf("Error fetching all logs: %v", err)
	}

	// Check if the created log exists in the list of all logs
	found := false
	for _, l := range allLogs {
		if logsAreEqual(&l, &log) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Created log not found in GetAllLogs result")
	}

	// Get logs by player ID
	playerLogs, err := repo.GetLogsByPlayerID(log.PlayerID)
	if err != nil {
		t.Fatalf("Error fetching logs by player ID: %v", err)
	}

	// Check if the created log exists in the list of logs by player ID
	found = false
	for _, l := range playerLogs {
		if logsAreEqual(&l, &log) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Created log not found in GetLogsByPlayerID result")
	}

	// Get logs by action
	actionLogs, err := repo.GetLogsByAction(log.Action)
	if err != nil {
		t.Fatalf("Error fetching logs by action: %v", err)
	}

	// Check if the created log exists in the list of logs by action
	found = false
	for _, l := range actionLogs {
		if logsAreEqual(&l, &log) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Created log not found in GetLogsByAction result")
	}

	// Get logs by time range (current timestamp - 1 minute to current timestamp + 1 minute)
	startTime := log.Timestamp - 60
	endTime := log.Timestamp + 60
	timeRangeLogs, err := repo.GetLogsByTimeRange(startTime, endTime)
	if err != nil {
		t.Fatalf("Error fetching logs by time range: %v", err)
	}

	// Check if the created log exists in the list of logs by time range
	found = false
	for _, l := range timeRangeLogs {
		if logsAreEqual(&l, &log) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Created log not found in GetLogsByTimeRange result")
	}

	// Delete the log
	err = repo.DeleteLog(id)
	if err != nil {
		t.Fatalf("Error deleting log: %v", err)
	}

	// Verify log deletion by trying to fetch it again
	_, err = repo.GetLogByID(id)
	if err == nil {
		t.Errorf("Expected log to be deleted, but it still exists")
	} else if !errors.Is(err, ErrLogNotFound) {
		t.Errorf("Expected ErrLogNotFound, got %v", err)
	}
}

// logsAreEqual checks if two logs are equal considering their fields.
func logsAreEqual(l1, l2 *models.Log) bool {
	if l1 == nil || l2 == nil {
		return false
	}
	return l1.ID == l2.ID &&
		l1.PlayerID == l2.PlayerID &&
		l1.Action == l2.Action &&
		l1.Timestamp == l2.Timestamp &&
		l1.CreatedAt == l2.CreatedAt &&
		l1.UpdatedAt == l2.UpdatedAt
}
