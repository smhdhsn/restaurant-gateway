package service

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"

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
func (s *EdibleMenuServ) List() (model.MenuItemListDTO, error) {
	return s.repo.List()
}
