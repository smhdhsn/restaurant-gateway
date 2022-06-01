package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/model"

	erpb "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/recipe"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// EdibleRecipeRepo contains repository's database connection.
type EdibleRecipeRepo struct {
	erc erpb.EdibleRecipeServiceClient
	ctx *context.Context
}

// NewEdibleMenuReository creates an instance of the remote repository with gRPC connection.
func NewEdibleRecipeRepository(ctx *context.Context, conn erpb.EdibleRecipeServiceClient) repositoryContract.EdibleRecipeRepository {
	return &EdibleRecipeRepo{
		erc: conn,
		ctx: ctx,
	}
}

// Store is responsible for storing recipe inside database.
func (s *EdibleRecipeRepo) Store(iListDTO model.MenuItemListDTO) error {
	rList := make([]*erpb.Recipe, len(iListDTO))
	for i, r := range iListDTO {
		rList[i] = &erpb.Recipe{
			FoodTitle:       r.Title,
			ComponentTitles: r.IngredientTitleList,
		}
	}

	req := erpb.RecipeStoreRequest{
		Recipes: rList,
	}

	_, err := s.erc.Store(*s.ctx, &req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.AlreadyExists:
				return repositoryContract.ErrDuplicateEntry
			case codes.Unknown:
				return repositoryContract.ErrUncaught
			}
		}

		return errors.Wrap(err, "error on calling store on edible gRPC server")
	}

	return nil
}
