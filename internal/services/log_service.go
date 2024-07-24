package services

import (
	"time"

	"oxo_game/internal/models"
	"oxo_game/internal/repositories"
	"sync"
)

type LogService interface {
	GetAllLogs() ([]models.Log, error)
	GetLogByID(id int) (*models.Log, error)
	CreateLog(log models.Log) (int, error)
	GetLogsByPlayerID(playerID int) ([]models.Log, error)
	GetLogsByAction(action string) ([]models.Log, error)
	GetLogsByTimeRange(startTime, endTime int64) ([]models.Log, error)
	DeleteLog(id int) error
}

type logService struct {
	logRepo repositories.LogRepository
	mu      sync.RWMutex
}

func NewLogService(repo repositories.LogRepository) LogService {
	return &logService{
		logRepo: repo,
	}
}

func (s *logService) GetAllLogs() ([]models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.logRepo.GetAllLogs()
}

func (s *logService) GetLogByID(id int) (*models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.logRepo.GetLogByID(id)
}

func (s *logService) CreateLog(log models.Log) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Add timestamps
	now := time.Now().Unix()
	log.Timestamp = now
	log.CreatedAt = now
	log.UpdatedAt = now

	return s.logRepo.CreateLog(log)
}

func (s *logService) GetLogsByPlayerID(playerID int) ([]models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.logRepo.GetLogsByPlayerID(playerID)
}

func (s *logService) GetLogsByAction(action string) ([]models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.logRepo.GetLogsByAction(action)
}

func (s *logService) GetLogsByTimeRange(startTime, endTime int64) ([]models.Log, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.logRepo.GetLogsByTimeRange(startTime, endTime)
}

func (s *logService) DeleteLog(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.logRepo.DeleteLog(id)
}
