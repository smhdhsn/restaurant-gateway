package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"

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
func (s *EdibleMenuServ) List() ([]*dto.Menu, error) {
	mListEntity, err := s.repo.List()
	if err != nil {
		return nil, errors.Wrap(err, "error on calling list on menu repository")
	}

	mListDTO := multipleMenuEntityToDTO(mListEntity)

	return mListDTO, nil
}

// multipleMenuEntityToDTO is responsible for transforming a list of menu entity into a list of menu dto struct.
func multipleMenuEntityToDTO(mListEntity []*entity.Menu) []*dto.Menu {
	mListDTO := make([]*dto.Menu, len(mListEntity))

	for i, mEntity := range mListEntity {
		iListDTO := make([]string, len(mEntity.Ingredients))

		copy(iListDTO, mEntity.Ingredients)

		mListDTO[i] = &dto.Menu{
			ID:          mEntity.ID,
			Title:       mEntity.Title,
			Ingredients: iListDTO,
		}
	}

	return mListDTO
}
