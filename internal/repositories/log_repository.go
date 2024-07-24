package repositories

import (
	"errors"
	"sync"

	"oxo_game/internal/models"
)

var (
	ErrLogNotFound = errors.New("log not found")
)

// LogRepository is the interface that wraps the basic CRUD operations for logs.
type LogRepository interface {
	GetAllLogs() ([]models.Log, error)
	GetLogByID(id int) (*models.Log, error)
	CreateLog(log models.Log) (int, error)
	GetLogsByPlayerID(playerID int) ([]models.Log, error)
	GetLogsByAction(action string) ([]models.Log, error)
	GetLogsByTimeRange(startTime, endTime int64) ([]models.Log, error)
	DeleteLog(id int) error
}

// InMemoryLogRepository is an example of a repository using in-memory storage.
type InMemoryLogRepository struct {
	mu     sync.RWMutex
	logs   map[int]models.Log
	autoID int
}

// NewInMemoryLogRepository creates a new InMemoryLogRepository.
func NewInMemoryLogRepository() *InMemoryLogRepository {
	return &InMemoryLogRepository{
		logs:   make(map[int]models.Log),
		autoID: 0,
	}
}

// GetAllLogs returns all logs.
func (r *InMemoryLogRepository) GetAllLogs() ([]models.Log, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	logs := make([]models.Log, 0, len(r.logs))
	for _, log := range r.logs {
		logs = append(logs, log)
	}
	return logs, nil
}

// GetLogByID returns the log with the given ID.
func (r *InMemoryLogRepository) GetLogByID(id int) (*models.Log, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	log, ok := r.logs[id]
	if !ok {
		return nil, ErrLogNotFound
	}
	return &log, nil
}

// CreateLog adds a new log and returns the new log's ID.
func (r *InMemoryLogRepository) CreateLog(log models.Log) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.autoID++
	log.ID = r.autoID
	r.logs[log.ID] = log
	return log.ID, nil
}

// GetLogsByPlayerID returns logs for a specific player ID.
func (r *InMemoryLogRepository) GetLogsByPlayerID(playerID int) ([]models.Log, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var logs []models.Log
	for _, log := range r.logs {
		if log.PlayerID == playerID {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

// GetLogsByAction returns logs for a specific action.
func (r *InMemoryLogRepository) GetLogsByAction(action string) ([]models.Log, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var logs []models.Log
	for _, log := range r.logs {
		if log.Action == action {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

// GetLogsByTimeRange returns logs within a specific time range.
func (r *InMemoryLogRepository) GetLogsByTimeRange(startTime, endTime int64) ([]models.Log, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var logs []models.Log
	for _, log := range r.logs {
		if log.Timestamp >= startTime && log.Timestamp <= endTime {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

// DeleteLog deletes the log with the given ID.
func (r *InMemoryLogRepository) DeleteLog(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.logs[id]; !ok {
		return ErrLogNotFound
	}
	delete(r.logs, id)
	return nil
}
