package service

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// EdibleRecipeServ contains repositories that will be used within this service.
type EdibleRecipeServ struct {
	repo repositoryContract.EdibleRecipeRepository
}

// NewEdibleRecipeService creates a edible's inventory service with it's dependencies.
func NewEdibleRecipeService(r repositoryContract.EdibleRecipeRepository) serviceContract.EdibleRecipeService {
	return &EdibleRecipeServ{
		repo: r,
	}
}

// Store is responsible for storing recipe inside database.
func (s *EdibleRecipeServ) Store(iList model.MenuItemListDTO) error {
	return s.repo.Store(iList)
}
