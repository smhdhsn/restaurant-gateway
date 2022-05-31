package service

import (
	"time"

	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// EdibleInventoryServ contains repositories that will be used within this service.
type EdibleInventoryServ struct {
	repo repositoryContract.EdibleInventoryRepository
}

// NewEdibleInventoryService creates a edible's inventory service with it's dependencies.
func NewEdibleInventoryService(r repositoryContract.EdibleInventoryRepository) serviceContract.EdibleInventoryService {
	return &EdibleInventoryServ{
		repo: r,
	}
}

// Buy is responsible for increasing stocks of a missing components.
func (s *EdibleInventoryServ) Buy(amount uint32, expiresAt time.Time) error {
	return s.repo.Buy(amount, expiresAt)
}

// Recycle is responsible for recycling expired or/and finished inventory stocks.
func (s *EdibleInventoryServ) Recycle(finished, expired bool) error {
	return s.repo.Recycle(finished, expired)
}

// Use is responsible for decreasing stocks of a component.
func (s *EdibleInventoryServ) Use(foodID uint32) error {
	return s.repo.Use(foodID)
}
