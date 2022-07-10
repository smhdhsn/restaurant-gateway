package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"

	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// OrderSubmissionServ contains repositories that will be used within this service.
type OrderSubmissionServ struct {
	repo repositoryContract.OrderSubmissionRepository
}

// NewOrderSubmissionService creates a order's submission service with it's dependencies.
func NewOrderSubmissionService(r repositoryContract.OrderSubmissionRepository) serviceContract.OrderSubmitService {
	return &OrderSubmissionServ{
		repo: r,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *OrderSubmissionServ) Submit(oDTO *dto.Order) error {
	oEntity := singleOrderDTOToEntity(oDTO)

	err := s.repo.Submit(oEntity)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrLackOfComponents) {
			return serviceContract.ErrLackOfComponents
		}

		return errors.Wrap(err, "error on calling submit on submission repository")
	}

	return nil
}

// singleOrderDTOToEntity is responsible for transforming an order dto to order entity struct.
func singleOrderDTOToEntity(oDTO *dto.Order) *entity.Order {
	return &entity.Order{
		FoodID: oDTO.FoodID,
		UserID: oDTO.UserID,
	}
}
