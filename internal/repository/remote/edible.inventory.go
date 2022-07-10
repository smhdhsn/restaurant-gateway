package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"

	inventoryProto "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/inventory"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// EdibleInventoryRepo contains repository's database connection.
type EdibleInventoryRepo struct {
	client inventoryProto.EdibleInventoryServiceClient
	ctx    context.Context
}

// NewEdibleMenuReository creates an instance of the remote repository with gRPC connection.
func NewEdibleInventoryRepository(ctx context.Context, conn inventoryProto.EdibleInventoryServiceClient) repositoryContract.EdibleInventoryRepository {
	return &EdibleInventoryRepo{
		client: conn,
		ctx:    ctx,
	}
}

// Recycle is responsible for recycling expired or/and finished inventory stocks.
func (r *EdibleInventoryRepo) Recycle(rEntity *entity.Recycle) error {
	req := singleRecycleEntityToReq(rEntity)

	_, err := r.client.Recycle(r.ctx, req)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			switch status.Code() {
			case codes.Internal:
				return errors.Wrap(err, "error inside edible gRPC server")
			}
		}

		return errors.Wrap(err, "error on calling recycle on edible gRPC server")
	}

	return nil
}

// singleRecycleEntityToReq is responsible for transforming a recycle entity into recycle request struct.
func singleRecycleEntityToReq(rEntity *entity.Recycle) *inventoryProto.InventoryRecycleRequest {
	return &inventoryProto.InventoryRecycleRequest{
		RecycleFinished: rEntity.Finished,
		RecycleExpired:  rEntity.Expired,
	}
}

// Buy is responsible for increasing stocks of a missing components.
func (r *EdibleInventoryRepo) Buy(bEntity *entity.Buy) error {
	req := singleBuyEntityToReq(bEntity)

	_, err := r.client.Buy(r.ctx, req)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			switch status.Code() {
			case codes.Internal:
				return errors.Wrap(err, "error inside edible gRPC server")
			}
		}

		return errors.Wrap(err, "error on calling buy on edible gRPC server")
	}

	return nil
}

// singleBuyEntityToReq is responsible for transforming a buy entity into buy request struct.
func singleBuyEntityToReq(bEntity *entity.Buy) *inventoryProto.InventoryBuyRequest {
	return &inventoryProto.InventoryBuyRequest{
		Amount:    bEntity.Amount,
		ExpiresAt: bEntity.ExpiresAt.Unix(),
	}
}
