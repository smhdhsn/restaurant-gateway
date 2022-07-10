package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"

	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// EdibleMenuServ contains repositories that will be used within this service.
type EdibleMenuServ struct {
	repo repositoryContract.EdibleMenuRepository
}

// NewEdibleMenuService creates a edible's inventory service with it's dependencies.
func NewEdibleMenuService(r repositoryContract.EdibleMenuRepository) serviceContract.EdibleMenuService {
	return &EdibleMenuServ{
		repo: r,
	}
}

// List returns a list of available items to order.
func (s *EdibleMenuServ) List() ([]*entity.Menu, error) {
	mListEntity, err := s.repo.List()
	if err != nil {
		return nil, errors.Wrap(err, "error on calling list on menu repository")
	}

	return mListEntity, nil
}
