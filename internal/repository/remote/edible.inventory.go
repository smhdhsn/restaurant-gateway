package remote

import (
	"context"
	"time"

	"github.com/pkg/errors"

	eipb "github.com/smhdhsn/restaurant-gateway/internal/protos/edible/inventory"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// EdibleInventoryRepo contains repository's database connection.
type EdibleInventoryRepo struct {
	eic eipb.EdibleInventoryServiceClient
	ctx *context.Context
}

// NewEdibleMenuReository creates an instance of the remote repository with gRPC connection.
func NewEdibleInventoryRepository(ctx *context.Context, conn eipb.EdibleInventoryServiceClient) repositoryContract.EdibleInventoryRepository {
	return &EdibleInventoryRepo{
		eic: conn,
		ctx: ctx,
	}
}

// Buy is responsible for increasing stocks of a missing components.
func (s *EdibleInventoryRepo) Buy(amount uint32, expiresAt time.Time) error {
	req := eipb.InventoryBuyRequest{
		Amount:    amount,
		ExpiresAt: expiresAt.Unix(),
	}

	_, err := s.eic.Buy(*s.ctx, &req)
	if err != nil {
		return errors.Wrap(err, "error on calling buy on edible gRPC server")
	}

	return nil
}

// Recycle is responsible for recycling expired or/and finished inventory stocks.
func (s *EdibleInventoryRepo) Recycle(finished, expired bool) error {
	req := eipb.InventoryRecycleRequest{
		RecycleFinished: finished,
		RecycleExpired:  expired,
	}

	_, err := s.eic.Recycle(*s.ctx, &req)
	if err != nil {
		return errors.Wrap(err, "error on calling recycle on edible gRPC server")
	}

	return nil
}

// Use is responsible for decreasing stocks of a component.
func (s *EdibleInventoryRepo) Use(foodID uint32) error {
	req := eipb.InventoryUseRequest{
		FoodId: foodID,
	}

	_, err := s.eic.Use(*s.ctx, &req)
	if err != nil {
		return errors.Wrap(err, "error on calling use on edible gRPC server")
	}

	return nil
}
