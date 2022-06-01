package service

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// OrderSubmitServ contains repositories that will be used within this service.
type OrderSubmitServ struct {
	repo repositoryContract.OrderSubmissionRepository
}

// NewOrderSubmitService creates a order's submission service with it's dependencies.
func NewOrderSubmitService(r repositoryContract.OrderSubmissionRepository) serviceContract.OrderSubmitService {
	return &OrderSubmitServ{
		repo: r,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *OrderSubmitServ) Submit(o *model.OrderDTO) error {
	return s.repo.Submit(o)
}
