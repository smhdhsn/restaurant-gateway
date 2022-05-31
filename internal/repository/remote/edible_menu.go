package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/model"

	empb "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/menu"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// EdibleMenuRepo contains repository's database connection.
type EdibleMenuRepo struct {
	emc empb.EdibleMenuServiceClient
	ctx *context.Context
}

// NewEdibleMenuReository creates an instance of the remote repository with gRPC connection.
func NewEdibleMenuRepository(ctx *context.Context, conn empb.EdibleMenuServiceClient) repositoryContract.EdibleMenuRepository {
	return &EdibleMenuRepo{
		emc: conn,
		ctx: ctx,
	}
}

// List returns a list of available items to order.
func (s *EdibleMenuRepo) List() (model.MenuItemListDTO, error) {
	resp, err := s.emc.List(*s.ctx, new(empb.MenuListRequest))
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				return nil, repositoryContract.ErrEmptyResult
			case codes.Unknown:
				return nil, repositoryContract.ErrUncaught
			}
		}

		return nil, errors.Wrap(err, "error on calling list on edible gRPC server")
	}

	iListDTO := make(model.MenuItemListDTO, len(resp.Foods))
	for i, f := range resp.GetFoods() {
		cTitleList := make([]string, len(f.GetIngredients()))
		for j, c := range f.GetIngredients() {
			cTitleList[j] = c.GetTitle()
		}

		iListDTO[i] = &model.MenuItemDTO{
			ID:                  f.GetId(),
			Title:               f.GetTitle(),
			IngredientTitleList: cTitleList,
		}
	}

	return iListDTO, nil
}
