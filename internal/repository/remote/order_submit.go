package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/model"

	ospb "github.com/smhdhsn/restaurant-gateway/internal/protos/order/submission"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// OrderSubmitRepo contains repository's database connection.
type OrderSubmitRepo struct {
	osr ospb.OrderSubmissionServiceClient
	ctx *context.Context
}

// NewOrderSubmitRepository creates an instance of the remote repository with gRPC connection.
func NewOrderSubmitRepository(ctx *context.Context, conn ospb.OrderSubmissionServiceClient) repositoryContract.OrderSubmissionRepository {
	return &OrderSubmitRepo{
		osr: conn,
		ctx: ctx,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (r *OrderSubmitRepo) Submit(o *model.OrderDTO) error {
	req := ospb.OrderSubmitRequest{
		FoodId: o.FoodID,
		UserId: o.UserID,
	}

	_, err := r.osr.Submit(*r.ctx, &req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.Unknown:
				return repositoryContract.ErrUncaught
			}
		}

		return errors.Wrap(err, "error on calling destroy on user gRPC server")
	}

	return nil
}
