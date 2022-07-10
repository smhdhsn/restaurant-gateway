package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"

	menuProto "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/menu"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// EdibleMenuRepo contains repository's database connection.
type EdibleMenuRepo struct {
	client menuProto.EdibleMenuServiceClient
	ctx    context.Context
}

// NewEdibleMenuReository creates an instance of the remote repository with gRPC connection.
func NewEdibleMenuRepository(ctx context.Context, conn menuProto.EdibleMenuServiceClient) repositoryContract.EdibleMenuRepository {
	return &EdibleMenuRepo{
		client: conn,
		ctx:    ctx,
	}
}

// List returns a list of available items to order.
func (r *EdibleMenuRepo) List() ([]*entity.Menu, error) {
	resp, err := r.client.List(r.ctx, new(menuProto.MenuListRequest))
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.Internal:
				return nil, errors.Wrap(err, "error inside edible gRPC server")
			}
		}

		return nil, errors.Wrap(err, "error on calling list on edible gRPC server")
	}

	mListEntity := multipleMenuRespToEntity(resp)

	return mListEntity, nil
}

// multipleMenuRespToEntity is responsible for transforming a list of menu response into a list of menu entity struct.
func multipleMenuRespToEntity(resp *menuProto.MenuListResponse) []*entity.Menu {
	mListEntity := make([]*entity.Menu, len(resp.Foods))

	for i, fResp := range resp.GetFoods() {
		cListResp := make([]string, len(fResp.GetIngredients()))

		for j, cResp := range fResp.GetIngredients() {
			cListResp[j] = cResp.GetTitle()
		}

		mListEntity[i] = &entity.Menu{
			ID:          fResp.GetId(),
			Title:       fResp.GetTitle(),
			Ingredients: cListResp,
		}
	}

	return mListEntity
}
