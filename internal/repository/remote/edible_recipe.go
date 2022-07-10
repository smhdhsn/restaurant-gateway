package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"

	recipeProto "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/recipe"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// EdibleRecipeRepo contains repository's database connection.
type EdibleRecipeRepo struct {
	erc recipeProto.EdibleRecipeServiceClient
	ctx context.Context
}

// NewEdibleMenuReository creates an instance of the remote repository with gRPC connection.
func NewEdibleRecipeRepository(ctx context.Context, conn recipeProto.EdibleRecipeServiceClient) repositoryContract.EdibleRecipeRepository {
	return &EdibleRecipeRepo{
		erc: conn,
		ctx: ctx,
	}
}

// Store is responsible for storing recipe inside database.
func (r *EdibleRecipeRepo) Store(rListEntity []*entity.Recipe) error {
	req := multipleRecipeEntityToStoreReq(rListEntity)

	_, err := r.erc.Store(r.ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.Internal:
				return errors.Wrap(err, "error inside edible gRPC server")
			}
		}

		return errors.Wrap(err, "error on calling store on edible gRPC server")
	}

	return nil
}

// multipleRecipeEntityToStoreReq is responsible for transforming a list of recipe entity into recipe request struct.
func multipleRecipeEntityToStoreReq(rListEntity []*entity.Recipe) *recipeProto.RecipeStoreRequest {
	rListReq := make([]*recipeProto.Recipe, len(rListEntity))

	for i, rEntity := range rListEntity {
		rListReq[i] = &recipeProto.Recipe{
			FoodTitle:       rEntity.Title,
			ComponentTitles: rEntity.Ingredients,
		}
	}

	return &recipeProto.RecipeStoreRequest{
		Recipes: rListReq,
	}
}
