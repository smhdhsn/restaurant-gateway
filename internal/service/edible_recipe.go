package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"

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
func (s *EdibleRecipeServ) Store(rListDTO []*dto.Recipe) error {
	rListEntity := multipleRecipeDTOToEntity(rListDTO)

	err := s.repo.Store(rListEntity)
	if err != nil {
		return errors.Wrap(err, "error on calling store on recipe repository")
	}

	return nil
}

// multipleRecipeDTOToEntity is responsible for transforming a list of recipe dto into a list of recipe entity struct.
func multipleRecipeDTOToEntity(rListDTO []*dto.Recipe) []*entity.Recipe {
	rListEntity := make([]*entity.Recipe, len(rListDTO))

	for i, rDTO := range rListDTO {
		rListEntity[i] = &entity.Recipe{
			Title:       rDTO.Title,
			Ingredients: rDTO.Ingredients,
		}
	}

	return rListEntity
}
