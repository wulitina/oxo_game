package repositories

import (
	"errors"
	"sync"
	"time"

	"oxo_game/internal/models"
)

type PaymentRepository interface {
	Create(payment *models.Payment) (int, error)
	GetById(id int) (*models.Payment, error)
}

type InMemoryPaymentRepository struct {
	mu       sync.RWMutex
	payments map[int]*models.Payment
	autoID   int
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{
		payments: make(map[int]*models.Payment),
		autoID:   0,
	}
}

func (r *InMemoryPaymentRepository) Create(payment *models.Payment) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.autoID++
	payment.ID = r.autoID
	payment.CreatedAt = time.Now().Unix()
	r.payments[payment.ID] = payment
	return payment.ID, nil
}

func (r *InMemoryPaymentRepository) GetById(id int) (*models.Payment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	payment, ok := r.payments[id]
	if !ok {
		return nil, errors.New("payment not found")
	}
	return payment, nil
}
