package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"

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

// Recycle is responsible for recycling expired or/and finished inventory stocks.
func (s *EdibleInventoryServ) Recycle(rDTO *dto.Recycle) error {
	rEntity := singleRecycleDTOToEntity(rDTO)

	err := s.repo.Recycle(rEntity)
	if err != nil {
		return errors.Wrap(err, "error on calling recycle on inventory repository")
	}

	return nil
}

// singleRecycleDTOToEntity is responsible for transforming a recycle dto into recycle entity struct.
func singleRecycleDTOToEntity(rDTO *dto.Recycle) *entity.Recycle {
	return &entity.Recycle{
		Expired:  rDTO.Expired,
		Finished: rDTO.Finished,
	}
}

// Buy is responsible for increasing stocks of a missing components.
func (s *EdibleInventoryServ) Buy(bDTO *dto.Buy) error {
	bEntity := singleBuyDTOToEntity(bDTO)

	err := s.repo.Buy(bEntity)
	if err != nil {
		return errors.Wrap(err, "error on calling buy on inventory repository")
	}

	return nil
}

// singleBuyDTOToEntity is responsible for transforming a buy dto into buy entity struct.
func singleBuyDTOToEntity(bDTO *dto.Buy) *entity.Buy {
	return &entity.Buy{
		Amount:    bDTO.Amount,
		ExpiresAt: bDTO.ExpiresAt,
	}
}
